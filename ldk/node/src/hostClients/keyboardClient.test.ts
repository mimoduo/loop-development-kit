import { mocked } from 'ts-jest/utils';
import * as Services from '../grpc/keyboard_grpc_pb';
import * as Messages from '../grpc/keyboard_pb';
import { ConnInfo } from '../grpc/broker_pb';
import KeyboardClient from './keyboardClient';
import { Session } from '../grpc/session_pb';
import { Logger } from '../logging';
import {
  captureMockArgument,
  createCallbackHandler,
  createEmptyStream,
  createStreamingHandler,
  identityCallback,
} from '../jest.helpers';
import { KeyboardHotkey } from '../grpc/keyboard_pb';

jest.mock('../grpc/keyboard_grpc_pb');

const hostClient = mocked(Services.KeyboardClient);

const logger = new Logger('test-logger');

describe('KeyboardClient', () => {
  let subject: KeyboardClient;
  let connInfo: ConnInfo.AsObject;
  let session: Session.AsObject;
  let waitForReadyMock: jest.Mock;
  let streamHotKeyMock: jest.Mock;
  let streamTextMock: jest.Mock;
  let streamCharMock: jest.Mock;
  let streamScanCodeMock: jest.Mock;

  beforeEach(() => {
    jest.resetAllMocks();
    subject = new KeyboardClient();
    connInfo = {
      address: 'a',
      serviceId: 1,
      network: 'n',
    };
    session = {
      loopid: 'LOOP_ID',
      token: 'TOKEN',
    };
    waitForReadyMock = jest.fn().mockImplementation(createCallbackHandler());

    hostClient.mockImplementation(() => {
      return {
        waitForReady: waitForReadyMock,
        keyboardHotkeyStream: streamHotKeyMock,
        keyboardTextStream: streamTextMock,
        keyboardCharacterStream: streamCharMock,
        keyboardScancodeStream: streamScanCodeMock,
      } as any;
    });
  });

  describe('#connect', () => {
    it('instantiates a new host client and waits for it to be ready', async () => {
      await expect(
        subject.connect(connInfo, session, logger),
      ).resolves.toBeUndefined();
    });
  });

  describe('#streamHotKey', () => {
    const key = 'a';
    const modifiers = { ctrlL: true };
    const convertedModifiers = 8;
    let sentRequest: Messages.KeyboardHotkeyStreamRequest;

    beforeEach(async () => {
      streamHotKeyMock = jest.fn().mockImplementation(createStreamingHandler());

      await subject.connect(connInfo, session, logger);

      subject.streamHotKey({ key, modifiers }, identityCallback);

      sentRequest = captureMockArgument(streamHotKeyMock);
    });

    it('should have configured the request with the right key and modifiers', () => {
      const hotKey = new KeyboardHotkey()
        .setKey(key)
        .setModifiers(convertedModifiers);

      expect(sentRequest.getHotkey()).toStrictEqual(hotKey);
    });

    it('should have attached the initial connection session to the request', () => {
      expect(sentRequest.getSession()?.toObject()).toStrictEqual(session);
    });
  });

  describe('#streamText', () => {
    let streamCallback: jest.Mock;
    let sentRequest: Messages.KeyboardTextStreamRequest;
    let sentResponse: Messages.KeyboardTextStreamResponse;

    beforeEach(async () => {
      sentResponse = new Messages.KeyboardTextStreamResponse().setText('hello');
      const stream = createEmptyStream();

      streamCallback = jest.fn().mockImplementation(identityCallback);
      streamTextMock = jest
        .fn()
        .mockImplementation(createStreamingHandler(stream));

      await subject.connect(connInfo, session, logger);

      subject.streamText(streamCallback);

      sentRequest = captureMockArgument(streamTextMock);
      stream.emit('data', sentResponse);
    });

    it('should stream the text back to the callback', () => {
      expect(streamCallback).toHaveBeenCalledWith(null, sentResponse.getText());
    });

    it('should have attached the initial connection session to the request', () => {
      expect(sentRequest.getSession()?.toObject()).toStrictEqual(session);
    });
  });

  describe('#streamChar', () => {
    let sentRequest: Messages.KeyboardCharacterStreamRequest;

    beforeEach(async () => {
      streamCharMock = jest.fn().mockImplementation(createStreamingHandler());

      await subject.connect(connInfo, session, logger);

      subject.streamChar(identityCallback);

      sentRequest = captureMockArgument(streamCharMock);
    });

    it('should have attached the initial connection session to the request', () => {
      expect(sentRequest.getSession()?.toObject()).toStrictEqual(session);
    });
  });

  describe('#streamScanCode', () => {
    let sentRequest: Messages.KeyboardScancodeStreamRequest;

    beforeEach(async () => {
      streamScanCodeMock = jest
        .fn()
        .mockImplementation(createStreamingHandler());

      await subject.connect(connInfo, session, logger);

      subject.streamScanCode(identityCallback);

      sentRequest = captureMockArgument(streamScanCodeMock);
    });

    it('should have attached the initial connection session to the request', () => {
      expect(sentRequest.getSession()?.toObject()).toStrictEqual(session);
    });
  });
});
