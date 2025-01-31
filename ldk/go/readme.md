# Loop Development Kit (LDK) for Go

> **PLEASE NOTE**: In order to provide a highly controlled and secure environment for our Loops, we will be transitioning to a JavaScript-only model and removing support for the C# and Go LDKs. This will also enable us to build additional features and Aptitudes more quickly. Please feel free to reach out to your Olive Helps developer contact if you have any questions.

## Installation

```shell
go get -u github.com/open-olive/loop-development-kit/ldk/go/v2
```

## Usage

Full documentation available on [pkg.go.dev](https://pkg.go.dev/github.com/open-olive/loop-development-kit/ldk/go/v2).

### Loops

Loops receive events and use them to generate relevant whispers. Loops choose which events they want to use and which they want to ignore.

This LDK can be used to write loops for Olive Helps. An example of the loop implementation is available [here](docs/loops.md).

#### Interface

Writing a loop boils down to writing an implementation for the `Loop` interface and serving it via `ldk.ServeLoopPlugin`.

```go
type Loop interface {
	LoopStart(Sidekick) error
	LoopStop() error
}
```

```go
func main() {
	// Instantiating Logger
	logger := ldk.NewLogger("demo-loop-go")

	// Instantiating Loop
	var loop ldk.Loop
	loop = &Loop{
		logger:         logger
	}

	ldk.ServeLoopPlugin(logger, loop) // ServeLoopPlugin, gives an ability to provide your own logger
}
```

**LoopStart** - The Loop should wait to start operating until this is called. The provided `Sidekick` reference should be stored in memory for continued use.

**LoopStop** - The Loop should stop operating when this is called.

#### Subscribing to Aptitudes

Inside `LoopStart`, you can subscribe to various aptitudes. Here's an example of subscribing to a couple:

```go
func (c *Loop) LoopStart(sidekick ldk.Sidekick) error {
	// ...

	handler := func (text string, err error) {
		// Respond to aptitude event here...
	}

	// Indicates if clipboard activities received from internal Olive Help windows will be propagated to the handler
	includeOliveHelpTraffic := true
	// Clipboard Configuration which takes handler and Clipboard traffic trigger
	clipboardListenConfiguration := ClipboardListenConfiguration{Handler: handler, IncludeOliveHelpTraffic: includeOliveHelpTraffic}

	if err := sidekick.Clipboard().Listen(l.ctx, clipboardListenConfiguration); err != nil {
		return err
	}

	if err := sidekick.UI().ListenSearchbar(l.ctx, handler); err != nil {
		return err
	}
}
```

Loops do not need to emit whispers in a 1:1 relationship with events, and may not use events at all. They could use some events, keep a history of events, or emit whispers only when several conditions are met.

#### List of possible Aptitudes to subscribe to

```go
type Sidekick interface {
	Clipboard() ClipboardService
	Vault() VaultService
	Whisper() WhisperService
	Keyboard() KeyboardService
	Process() ProcessService
	Cursor() CursorService
	Filesystem() FilesystemService
	Window() WindowService
	UI() UIService
	Network() NetworkService
}
```

#### Lifecycle

1. Olive Helps executes plugin process.
1. Olive Helps calls `LoopStart`, sending the `Sidekick` reference to the plugin.
1. The loop subscribes to one or more aptitudes in `LoopStart`.
1. When the loop is notified of an aptitudes event, it processes it and calls the `Whisper` method on the `Sidekick` reference to emit a whisper.
1. On user disabling the loop, Olive Helps calls `LoopStop` then sends `sigterm` to the process.
1. On Olive Helps shutdown, Olive Helps calls `LoopStop` then sends `sigterm` to the process.

### Running Locally

#### Local Plugin Command (Recommended)

Olive Helps lets you add a local command as Local Plugins:

1. Open Olive Helps.
2. Open the Loop Library:
   1. Click the Hamburger icon.
   2. Click Loop Library.
3. Click the Install Local Plugin button:
4. Select the working directory for the command.
5. Enter the command to be executed, including any arguments.
6. Click Install.

The command will be installed as a plugin. If you need to change the command or its arguments you'll need remove it and then add the new commands. If you are developing locally, it can be helpful to give it the `go run` command, as this will remove the need to recompile each time to check your changes:

```sh
go run main.go
```

Be sure that when you selected the working directory in `Step 5` that it is the root of your project that has the `main.go` file located in it.

### Troubleshooting and Debugging

Olive Helps logs are available in the following directories for your OS:

```shell
~/Library/Logs/Olive\ Helps   # MacOS
/var/log/Olive\ Helps         # Linux
%AppData%\Olive Helps\Logs\   # Windows
```

#### Tailing Logs

It can be useful to tail the log as you develop to see any errors, warnings, or info messages. To do so, you can issue any of the following commands:

##### Linux

```sh
tail -F /var/log/Olive\ Helps/*.log
```

##### Mac

```sh
tail -F ~/Library/Logs/Olive\ Helps/*.log
```

##### Windows

From `Powershell` (not a cmd shell):

```powershell
Get-Content $env:AppData\Olive Helps\Logs\Olive Helps-2020.08.12-1dcc37a.log -Tail 10 –Wait
```

Note: File names may differ from the examples above.

## Development

### protoc

The LDK utilizes `protoc` to create the `protobuf` that defines the contract that is utilized by GRPC.

- Install protobuf with Homebrew `brew link protobuf`
- Make your changes to `proto/ldk.proto`
- Then generate the new protobuf file with `make proto/ldk.pb.go`

Since `proto/ldk.pb.go` is auto-generated, do not edit this file directly.
