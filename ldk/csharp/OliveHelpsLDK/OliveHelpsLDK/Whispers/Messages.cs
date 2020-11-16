using System;
using System.Collections.Immutable;
using Google.Protobuf.WellKnownTypes;
using Proto;

namespace OliveHelpsLDK.Whispers
{
    public struct WhisperStyle
    {
        public string BackgroundColor;
        public string PrimaryColor;
        public string HighlightColor;
    }

    public struct WhisperConfig
    {
        public string Icon;
        public string Label;
        public WhisperStyle Style;
    }

    public struct WhisperMarkdown
    {
        public string Markdown;
        public WhisperConfig Config;
    }

    public struct WhisperConfirm
    {
        public string Markdown;
        public WhisperConfig Config;
        public string ResolveLabel;
        public string RejectLabel;
    }

    public struct WhisperForm
    {
        public string Markdown;
        public WhisperConfig Config;
        public string SubmitLabel;
        public string CancelLabel;
        public ImmutableDictionary<string, Forms.Inputs.IBase<object>> Inputs;
    }
}

namespace OliveHelpsLDK.Whispers.Forms.Inputs
{
    public interface IBase<T>
    {
        string Label { get; }
        string Tooltip { get; }
        int Order { get; }

        T ToProto();
    }

    public interface ICheckbox : IBase<Proto.WhisperFormInput.Types.Checkbox>
    {
        bool Value { get; }
    }

    public interface IEmail : IBase<Proto.WhisperFormInput.Types.Email>
    {
        string Value { get; }
    }

    public interface IMarkdown : IBase<Proto.WhisperFormInput.Types.Markdown>
    {
        string Value { get; }
    }

    public interface INumber : IBase<Proto.WhisperFormInput.Types.Number>
    {
        float Value { get; }
        float Min { get; }
        float Max { get; }
    }

    public interface IPassword : IBase<Proto.WhisperFormInput.Types.Password>
    {
    }

    public interface IRadio : IBase<Proto.WhisperFormInput.Types.Radio>
    {
        string[] Options { get; }
    }

    public interface ISelect : IBase<Proto.WhisperFormInput.Types.Select>
    {
        string[] Options { get; }
    }

    public interface ITelephone : IBase<WhisperFormInput.Types.Tel>
    {
        string Value { get; }
        string Pattern { get; }
    }

    public interface IText : IBase<Proto.WhisperFormInput.Types.Text>
    {
        string Value { get; }
    }

    public interface ITime : IBase<Proto.WhisperFormInput.Types.Time>
    {
        DateTimeOffset Value { get; }
    }

    public struct Checkbox : ICheckbox
    {
        public string Label { get; }
        public string Tooltip { get; }
        public int Order { get; }
        public bool Value { get; }

        public WhisperFormInput.Types.Checkbox ToProto()
        {
            return new WhisperFormInput.Types.Checkbox
            {
                Label = Label,
                Tooltip = Tooltip,
                Order = checked((uint) Order),
                Value = Value
            };
        }
    }

    public struct Email : IEmail
    {
        public string Label { get; }
        public string Tooltip { get; }
        public int Order { get; }

        public string Value { get; }

        public WhisperFormInput.Types.Email ToProto()
        {
            return new WhisperFormInput.Types.Email
            {
                Label = Label,
                Tooltip = Tooltip,
                Order = checked((uint) Order),
                Value = Value
            };
        }
    }

    public struct Markdown : IMarkdown
    {
        public string Label { get; }
        public string Tooltip { get; }
        public int Order { get; }
        public string Value { get; }

        public WhisperFormInput.Types.Markdown ToProto()
        {
            return new WhisperFormInput.Types.Markdown
            {
                Label = Label,
                Tooltip = Tooltip,
                Order = checked((uint) Order),
                Value = Value
            };
        }
    }

    public struct Number : INumber
    {
        public string Label { get; }
        public string Tooltip { get; }
        public int Order { get; }

        public float Value { get; }

        public float Min { get; }

        public float Max { get; }

        public WhisperFormInput.Types.Number ToProto()
        {
            return new WhisperFormInput.Types.Number
            {
                Label = Label,
                Tooltip = Tooltip,
                Order = checked((uint) Order),
                Value = Value,
                Max = Max,
                Min = Min,
            };
        }
    }

    public struct Password : IPassword
    {
        public string Label { get; }
        public string Tooltip { get; }
        public int Order { get; }
        public WhisperFormInput.Types.Password ToProto()
        {
            return new WhisperFormInput.Types.Password
            {
                Label = Label,
                Order = checked((uint)Order),
                Tooltip = Tooltip
            };
        }
    }

    public struct Radio : IRadio
    {
        public string Label { get; }
        public string Tooltip { get; }
        public int Order { get; }
        public WhisperFormInput.Types.Radio ToProto()
        {
            var radio = new WhisperFormInput.Types.Radio
            {
                Label = Label,
                Order = checked((uint)Order),
                Tooltip = Tooltip,
            };
            foreach (var option in Options)
            {
                radio.Options.Add(option);
            }
            return radio;
        }

        public string[] Options { get; }
    }

    public struct Select : ISelect
    {
        public string Label { get; }
        public string Tooltip { get; }
        public int Order { get; }
        public WhisperFormInput.Types.Select ToProto()
        {
            var select = new WhisperFormInput.Types.Select
            {
                Label = Label,
                Order = checked((uint)Order),
                Tooltip = Tooltip,
            };
            foreach (var option in Options)
            {
                select.Options.Add(option);
            }
            return select;
        }

        public string[] Options { get; }
    }

    public struct Telephone : ITelephone
    {
        public string Label { get; }
        public string Tooltip { get; }
        public int Order { get; }

        public string Value { get; }

        public string Pattern { get; }

        public WhisperFormInput.Types.Tel ToProto()
        {
            return new WhisperFormInput.Types.Tel
            {
                Label = Label,
                Order = checked((uint)Order),
                Tooltip = Tooltip,
                Value = Value,
                Pattern = Pattern,
            };
        }
    }

    public struct Text : IText
    {
        public string Label { get; }
        public string Tooltip { get; }
        public int Order { get; }

        public string Value { get; }

        public WhisperFormInput.Types.Text ToProto()
        {
            return new WhisperFormInput.Types.Text
            {
                Label = Label,
                Order = checked((uint)Order),
                Tooltip = Tooltip,
                Value = Value,
            };
        }
    }

    public struct Time : ITime
    {
        public string Label { get; }
        public string Tooltip { get; }
        public int Order { get; }
        public DateTimeOffset Value { get; }

        public WhisperFormInput.Types.Time ToProto()
        {
            return new WhisperFormInput.Types.Time
            {
                Label = Label,
                Order = checked((uint) Order),
                Tooltip = Tooltip,
                Value = Timestamp.FromDateTimeOffset(Value),
            };
        }
    }
}