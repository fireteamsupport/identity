package email

import (
    "github.com/arturoguerra/go-logging"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ses"
)

var log = logging.New()

const (
    CharSet = "UTF-8"
)

type (
    email struct {
        Session *ses.SES
        Sender string
    }

    Email interface {
        Send(string, string, string)
    }
)

func NewDefault() (error, Email) {
    err, cfg := NewEnvConfig()
    if err != nil {
        return err, nil
    }

    err, instance := New(cfg)
    return err, instance
}

func New(cfg *Config) (error, Email) {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(cfg.Region),
    })

    if err != nil {
        log.Fatal(err)
    }

    svc := ses.New(sess)

    return nil, &email{
        Session: svc,
        Sender: cfg.Sender,
    }
}


func (em *email) Send(recipient, subject, body string) {
    input := &ses.SendEmailInput{
        Destination: &ses.Destination{
            CcAddresses: []*string{},
            ToAddresses: []*string{
                aws.String(recipient),
            },
        },
        Message: &ses.Message{
            Body: &ses.Body{
                Html: &ses.Content{
                    Charset: aws.String(CharSet),
                    Data:    aws.String(body),
                },
            },
            Subject: &ses.Content{
                Charset: aws.String(CharSet),
                Data:    aws.String(subject),
            },
        },
        Source: aws.String(em.Sender),
    }

    _, err := em.Session.SendEmail(input)

    if err != nil {
        log.Error(err)
    }
}
