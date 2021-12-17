import * as React from 'react';
import '@oliveai/ldk';
import * as ReactWhisper from '@oliveai/ldk/dist/whisper/react/renderer';
// Display all enums rather than importing with * for better referencing as a test loop
import {
  AlignItems,
  ButtonSize,
  ButtonStyle,
  Color,
  CustomHeight,
  DateTimeType,
  Direction,
  IconSize,
  JustifyContent,
  Urgency,
  MarkdownWhisperCopyMode,
  MessageWhisperCopyMode,
  OpenDirection,
  ProgressShape,
  StyleSize,
  TextAlign,
  WidthSize,
} from '@oliveai/ldk/dist/whisper/types';

const ComponentHotTub = () => {
  const externalLink = 'https://docs.oliveai.dev/';
  const options = [
    {
      label: 'Option One',
      value: 'optionone'
    },
    {
      label: 'Option Two',
      value: 'optiontwo'
    },
    {
      label: 'Option Three',
      value: 'optionthree'
    },
    {
      label: 'Option Four',
      value: 'optionfour'
    }
  ];

  const selectOptions = [
    "Option One",
    "Option Two",
    "Option Three",
  ];

  const onEvent = (error, value) => {
    console.log(`React Test Loop: ${value}${error ? `, ${error}` : ''}`)
  };

  return (
    <oh-whisper 
      label="React" 
      onClose={() => {}}
    >
      <oh-autocomplete
        freeSolo={false}
        label="Autocomplete"
        loading={false}
        multiple={false}
        onChange={(error, value) => onEvent(error, value)}
        onSelect={(error, value) => onEvent(error, value)}
        options={options}
        tooltip="Autocomplete"
        value=""
      />
      <oh-autocomplete
        freeSolo={true}
        label="Autocomplete Freesolo"
        loading={false}
        multiple={false}
        onChange={(error, value) => onEvent(error, value)}
        onSelect={(error, value) => onEvent(error, value)}
        options={options}
        tooltip="Autocomplete Freesolo"
        value=""
      />
      <oh-autocomplete
        freeSolo={false}
        label="Autocomplete Freesolo"
        loading={false}
        multiple={true}
        onChange={(error, value) => onEvent(error, value)}
        onSelect={(error, value) => onEvent(error, value)}
        options={options}
        tooltip="Autocomplete Multiple"
        value=""
      />
      <oh-box
        alignItems={AlignItems.Center}
        direction={Direction.Horizontal}
        justifyContent={JustifyContent.Center}
        onClick={(error, value) => onEvent(error, value)}
      >
        Align Center, Direction Horizontal, Justify Center
      </oh-box>
      {/* TODO: Breaks loop upon startup */}
      {/* <oh-breadcrumbs 
        links={[
          {
            href: {externalLink},
            text: 'Google',
            onClick: (error, value) => onEvent(error, value),
          },
          {
            href: 'apple.com',
            text: 'Apple Web Developer',
            onClick: (error, value) => onEvent(error, value),
          },
          {
            href: 'meta.com',
            text: 'Meta Oculus Program',
            onClick: (error, value) => onEvent(error, value),
          }
        ]}
      /> */}
      <oh-box
        alignItems={AlignItems.Center}
        direction={Direction.Horizontal}
        justifyContent={JustifyContent.Center}
      >
        <oh-button
          buttonStyle={ButtonStyle.Primary}
          disabled={false}
          label="Small"
          onClick={(error, value) => onEvent(error, value)}
          size={ButtonSize.Small}
          tooltip="Small"
        />
        <oh-button
          buttonStyle={ButtonStyle.Secondary}
          disabled={false}
          label="Large"
          onClick={(error, value) => onEvent(error, value)}
          size={ButtonSize.Large}
          tooltip="Large"
        />
        <oh-button
          buttonStyle={ButtonStyle.Text}
          disabled={false}
          label="Small Text"
          onClick={(error, value) => onEvent(error, value)}
          size={ButtonSize.Small}
          tooltip="Small Text"
        />
        <oh-button
          buttonStyle={ButtonStyle.Text}
          disabled={false}
          label="Large Text"
          onClick={(error, value) => onEvent(error, value)}
          size={ButtonSize.Large}
          tooltip="Large Text"
        />
      </oh-box>
      <oh-checkbox
        label="Checkbox One"
        onChange={(error, value) => onEvent(error, value)}
        tooltip="Checkbox One"
        value={false}
      />
      <oh-checkbox
        label="Checkbox Two"
        onChange={(error, value) => onEvent(error, value)}
        tooltip="Checkbox Two"
        value={true}
      />
      <oh-collapse-box
        label="CollapseBox Top"
        onClick={(error, value) => onEvent(error, value)}
        open={false}
        openDirection={OpenDirection.Top}
      >
        When a baby peperomia meets salami, there is often a dispute in who is the spiciest creature. 
        Peperomia, the queen of the forest floor, will strike a pose upon first glance. 
        Even a moment is enough to freeze the toughest taste buds in their tracks.
        Salami, the terminator of all who read blues, constructs pillows of agony with every step.
        With the slightest reverberation, an iron clad tongue has but no chance to break towards freedom.
        What do you do?
      </oh-collapse-box>
      <oh-collapse-box
        label="CollapseBox Bottom"
        onClick={(error, value) => onEvent(error, value)}
        open={true}
        openDirection={OpenDirection.Bottom}
      >
        Collapsed Text
      </oh-collapse-box>
      <oh-divider />
      <oh-drop-zone 
        accept={[
          'jpg', 
          'jpeg', 
          'png',
        ]}
        label="DropZone"
        limit={3}
        noun="Custom Noun"
        onDrop={(error, value) => onEvent(error, value)}
        tooltip="DropZone Tooltip"
      />
      <oh-drop-zone 
        label="DropZone Error Without Optional Fields "
        onDrop={(error, value) => onEvent(error, value)}
        tooltip="DropZone Tooltip"
        validationError="Validation Error"
      />
      <oh-box
        alignItems={AlignItems.Center}
        direction={Direction.Horizontal}
        justifyContent={JustifyContent.Center}
      >
        <oh-icon 
          color={Color.Black}
          name="settings"
          onClick={(error, value) => onEvent(error, value)}
          size={IconSize.Small}
          tooltip="Settings Icon"
        />
        <oh-icon 
          color={Color.Grey}
          name="touch_app"
          onClick={(error, value) => onEvent(error, value)}
          size={IconSize.Medium}
          tooltip="Touch App Icon"
        />
        <oh-icon 
          color={Color.White}
          name="task_alt"
          onClick={(error, value) => onEvent(error, value)}
          size={IconSize.Large}
          tooltip="Task Alt Icon"
        />
        <oh-icon 
          color={Color.WhisperStrip}
          name="delete"
          onClick={(error, value) => onEvent(error, value)}
          size={IconSize.XLarge}
          tooltip="Delete Icon"
        />
      </oh-box>
      <oh-link 
        href={externalLink}
        onClick={(error, value) => onEvent(error, value)}
        style={Urgency.None}
        text="Link None Left"
        textAlign={TextAlign.Left}
      />
      <oh-link 
        href={externalLink}
        onClick={(error, value) => onEvent(error, value)}
        style={Urgency.Error}
        text="Link Error Center"
        textAlign={TextAlign.Center}
      />
      <oh-link 
        href={externalLink}
        onClick={(error, value) => onEvent(error, value)}
        style={Urgency.Success}
        text="Link Success Right"
        textAlign={TextAlign.Right}
      />
      <oh-link 
        href={externalLink}
        onClick={(error, value) => onEvent(error, value)}
        style={Urgency.Warning}
        text="Link Warning"
      />
      <oh-list-pair
        copyable={false}
        labelCopyable={false}
        label="Label None Not Copyable"
        onCopy={(error, value) => onEvent(error, value)}
        style={Urgency.None}
        value="Value None Not Copyable"
      />
      <oh-list-pair
        copyable={true}
        labelCopyable={true}
        label="Label Error Copyable"
        onCopy={(error, value) => onEvent(error, value)}
        style={Urgency.Error}
        value="Value Error Copyable"
      />
      <oh-list-pair
        copyable={true}
        labelCopyable={false}
        label="Label Success Not Copyable"
        onCopy={(error, value) => onEvent(error, value)}
        style={Urgency.Success}
        value="Value Success Copyable"
      />
      <oh-list-pair
        copyable={false}
        labelCopyable={true}
        label="Label Warning Copyable"
        onCopy={(error, value) => onEvent(error, value)}
        style={Urgency.Warning}
        value="Value Warning Not Copyable"
      />
      <oh-markdown
        body="# Markdown Header 1
        ## Markdown Header 2"
        copyable={MarkdownWhisperCopyMode.Body}
        onCopy={(error, value) => onEvent(error, value)}
        onLinkClick={(error, value) => onEvent(error, value)}
        tooltip="Markdown Tooltip"
      />
      <oh-message
        copyable={MessageWhisperCopyMode.Header}
        body="Message Accent"
        header="Message Accent"
        onCopy={(error, value) => onEvent(error, value)}
        style={Color.Accent}
        tooltip="Message Accent Tooltip"
      />
      <oh-message
        copyable={MessageWhisperCopyMode.Header}
        body="Message Black"
        header="Message Black"
        onCopy={(error, value) => onEvent(error, value)}
        style={Color.Black}
        tooltip="Message Black Tooltip"
      />
      <oh-message
        copyable={MessageWhisperCopyMode.Body}
        body="Message Grey"
        header="Message Grey"
        onCopy={(error, value) => onEvent(error, value)}
        style={Color.Grey}
        tooltip="Message Grey Tooltip"
      />
      <oh-message
        body="Message WhisperStrip"
        header="Message WhisperStrip"
        onCopy={(error, value) => onEvent(error, value)}
        style={Color.WhisperStrip}
        tooltip="Message WhisperStrip Tooltip"
      />
      <oh-message
        body="Message None"
        header="Message None"
        onCopy={(error, value) => onEvent(error, value)}
        tooltip="Message None Tooltip"
      />
      <oh-message
        body="Message Error Left"
        header="Message Error Left"
        onCopy={(error, value) => onEvent(error, value)}
        style={Urgency.Error}
        textAlign={TextAlign.Left}
        tooltip="Message Error Left Tooltip"
      />
      <oh-message
        body="Message Success Center"
        header="Message Success Center"
        onCopy={(error, value) => onEvent(error, value)}
        style={Urgency.Success}
        textAlign={TextAlign.Center}
        tooltip="Message Success Center Tooltip"
      />
      <oh-message
        body="Message Warning Right"
        header="Message Warning Right"
        onCopy={(error, value) => onEvent(error, value)}
        style={Urgency.Warning}
        textAlign={TextAlign.Right}
        tooltip="Message Warning Right Tooltip"
      />
      <oh-box
        alignItems={AlignItems.Center}
        direction={Direction.Horizontal}
        justifyContent={JustifyContent.Center}
      >
        <oh-progress
          shape={ProgressShape.Circular}
          size={StyleSize.None}
        />
        <oh-progress
          determinate={20}
          shape={ProgressShape.Circular}
          size={StyleSize.None}
        />
        <oh-progress
          determinate={40}
          shape={ProgressShape.Circular}
          size={StyleSize.Small}
        />
        <oh-progress
          determinate={60}
          shape={ProgressShape.Circular}
          size={StyleSize.Medium}
        />
        <oh-progress
          determinate={80}
          shape={ProgressShape.Circular}
          size={StyleSize.Large}
        />
      </oh-box>
      <oh-progress
        shape={ProgressShape.Linear}
        size={StyleSize.None}
      />
      <oh-progress
        determinate={20}
        shape={ProgressShape.Linear}
        size={StyleSize.None}
      />
      <oh-radio-group
        onSelect={(error, value) => onEvent(error, value)}
        options={[
          "Radio One",
          "Radio Two",
          "Radio Three",
        ]}
        selected={0}
      />
      <oh-rich-text-editor
        onBlur={(error, value) => onEvent(error, value)}
        onChange={(error, value) => onEvent(error, value)}
        onFocus={(error, value) => onEvent(error, value)}
        tooltip="Rich Text Editor"
      />
      <oh-section-title
        body="Section Title None Left"
        textAlign={TextAlign.Left}
      />
      <oh-section-title
        backgroundStyle={Color.Grey}
        body="Section Title Grey Center"
        textAlign={TextAlign.Center}
      />
      <oh-section-title
        backgroundStyle={Color.White}
        body="Section Title White Right"
        textAlign={TextAlign.Right}
      />
      <oh-select 
        label="Select Label"
        loading={false}
        onChange={(error, value) => onEvent(error, value)}
        onSelect={(error, value) => onEvent(error, value)}
        options={selectOptions}
        tooltip="Select Tooltip"
        value=""
      />
      <oh-text-input 
        label="Text Input"
        tooltip="Text Input"
        onBlur={(error, value) => onEvent(error, value)}
        onFocus={(error, value) => onEvent(error, value)}
        onChange={(error, value) => onEvent(error, value)}
        validationError=""
        value=""
      />
      <oh-number 
        label="Number Input"
        max={200}
        min={-200}
        onBlur={(error, value) => onEvent(error, value)}
        onFocus={(error, value) => onEvent(error, value)}
        onChange={(error, value) => onEvent(error, value)}
        step={20}
        tooltip="Number Input"
        validationError=""
        value={0}
      />
      <oh-telephone 
        label="Telephone Input"
        onBlur={(error, value) => onEvent(error, value)}
        onFocus={(error, value) => onEvent(error, value)}
        onChange={(error, value) => onEvent(error, value)}
        tooltip="Telephone Input"
        validationError=""
        value=""
      />
      <oh-password 
        label="Password Input"
        onBlur={(error, value) => onEvent(error, value)}
        onFocus={(error, value) => onEvent(error, value)}
        onChange={(error, value) => onEvent(error, value)}
        tooltip="Password Input"
        validationError=""
        value=""
      />
      <oh-email 
        label="Email Input"
        onBlur={(error, value) => onEvent(error, value)}
        onFocus={(error, value) => onEvent(error, value)}
        onChange={(error, value) => onEvent(error, value)}
        tooltip="Email Input"
        validationError=""
        value=""
      />
      <oh-datetime 
        dateTimeType={DateTimeType.Date}
        label="Date Input"
        max={new Date()}
        min={new Date(0)}
        onBlur={(error, value) => onEvent(error, value)}
        onFocus={(error, value) => onEvent(error, value)}
        onChange={(error, value) => onEvent(error, value)}
        tooltip="Date Input"
        validationError=""
      />
      <oh-datetime 
        dateTimeType={DateTimeType.DateTime}
        label="DateTime Input"
        max={new Date()}
        min={new Date(0)}
        onBlur={(error, value) => onEvent(error, value)}
        onFocus={(error, value) => onEvent(error, value)}
        onChange={(error, value) => onEvent(error, value)}
        tooltip="DateTime Input"
        validationError=""
      />
      <oh-datetime 
        dateTimeType={DateTimeType.Time}
        label="Time Input"
        max={new Date()}
        min={new Date(0)}
        onBlur={(error, value) => onEvent(error, value)}
        onFocus={(error, value) => onEvent(error, value)}
        onChange={(error, value) => onEvent(error, value)}
        tooltip="Time Input"
        validationError=""
      />
    </oh-whisper>
  );
};

ReactWhisper.renderNewWhisper(<ComponentHotTub/>);