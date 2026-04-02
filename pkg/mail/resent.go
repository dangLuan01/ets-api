package mail

import (
	"context"
	"net/http"
	"strings"

	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/resend/resend-go/v2"
)

type ResentConfig struct {
	MailSender string
	NameSender string	
	ResentApiKey string
}

type ResentProvider struct {
	client *http.Client
	config *ResentConfig
}

func NewResentProvider(config *MailConfig) (EmailProviderService, error) {

	resentConfig, ok := config.ProviderConfig["resent"].(map[string]any)
	if !ok {
		return nil, utils.NewError(string(utils.ErrCodeInternal), "Invalid or missing Resent config.")
	}


	return &ResentProvider{
		client: &http.Client{Timeout: config.Timeout},
		config: &ResentConfig{
			MailSender: resentConfig["mail_sender"].(string),
			NameSender: resentConfig["name_sender"].(string),
			ResentApiKey: resentConfig["resent_api_key"].(string),
		},
	}, nil
}

func (p *ResentProvider) SendMail(ctx context.Context, email *Email) error {
	var htmlTemplate string
	data := map[string]string{}

	switch email.Category {
	case "otp":
		htmlTemplate = `
		<!DOCTYPE html>
		<html xmlns="http://www.w3.org/1999/xhtml">
		
		<head>
		<title></title>
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<style type="text/css">
			#outlook a {
			padding: 0;
			}
		
			.ReadMsgBody {
			width: 100%;
			}
		
			.ExternalClass {
			width: 100%;
			}
		
			.ExternalClass * {
			line-height: 100%;
			}
		
			body {
			margin: 0;
			padding: 0;
			-webkit-text-size-adjust: 100%;
			-ms-text-size-adjust: 100%;
			}
		
			table,
			td {
			border-collapse: collapse;
			mso-table-lspace: 0pt;
			mso-table-rspace: 0pt;
			}
		
		</style>
		<style type="text/css">
			@media only screen and (max-width:480px) {
			@-ms-viewport {
				width: 320px;
			}
			@viewport {
				width: 320px;
			}
			}
		</style>
		<link href="https://fonts.googleapis.com/css2?family=Open+Sans:wght@400;600&display=swap" rel="stylesheet" type="text/css">
		<style type="text/css">
			@import url('https://fonts.googleapis.com/css2?family=Open+Sans:wght@400;600&display=swap');
		</style>
		<style type="text/css">
			@media only screen and (max-width:595px) {
			.container {
				width: 100% !important;
			}
			.button {
				display: block !important;
				width: auto !important;
			}
			}
		</style>
		</head>
		
		<body style="font-family: 'Inter', sans-serif; background: #E5E5E5;">
		<table width="100%" cellspacing="0" cellpadding="0" border="0" align="center" bgcolor="#F6FAFB">
			<tbody>
			<tr>
				<td valign="top" align="center">
				<table class="container" width="600" cellspacing="0" cellpadding="0" border="0">
					<tbody>
					<tr>
						<td style="padding:48px 0 30px 0; text-align: center; font-size: 14px; color: #4C83EE;">
						<h2>XOAI LAC STREAMING</h2>
						</td>
					</tr>
					<tr>
						<td class="main-content" style="padding: 48px 30px 40px; color: #000000;" bgcolor="#ffffff">
						<table width="100%" cellspacing="0" cellpadding="0" border="0">
							<tbody>
							<tr>
								<td style="padding: 0 0 24px 0; font-size: 18px; line-height: 150%; font-weight: bold; color: #000000; letter-spacing: 0.01em;">
								Verification Code
								</td>
							</tr>
							<tr>
								<td style="padding: 0 0 10px 0; font-size: 14px; line-height: 150%; font-weight: 400; color: #000000; letter-spacing: 0.01em;">
								Here is the code you requested to your account <span style="color: #4C83EE;">{{user_email}}</span>. Enter this code on the verification page.
								</td>
							</tr>
							<tr>
								<td style="padding: 0 0 10px 0;height:24px;font-size:20px;font-weight:700;color:rgba(42,77,143,1);line-height:24px" align="left">
								{{pass_reset_link}}
							</td>
							</tr>
							<tr>
								<td style="padding:0 0 10px 0;font-size:12px;font-weight:500;color:rgba(153,153,153,1);line-height:22px" align="left">
								Your verification code is valid for 6 hours.
								<span style="color:rgba(70,77,98,1)"></span>
								</td>
							</tr>
							<tr>
								<td style="padding: 0 0 16px;">
								<span style="display: block; width: 117px; border-bottom: 1px solid #8B949F;"></span>
								</td>
							</tr>
							</tbody>
						</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 24px 0 48px; font-size: 0px;">
						<div class="outlook-group-fix" style="padding: 0 0 20px 0; vertical-align: top; display: inline-block; text-align: center; width:100%;">
							<span style="padding: 0; font-size: 11px; line-height: 15px; font-weight: normal; color: #8B949F;">Xoai Lac<br/>Company Physical Address</span>
						</div>
						</td>
					</tr>
					</tbody>
				</table>
				</td>
			</tr>
			</tbody>
		</table>
		</body>
		</html>
		`
		data = map[string]string{
			"user_email":      email.ToOfResent[0],
			"pass_reset_link": email.Html,
		}
	case "payment":
		htmlTemplate = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title></title>
			<style>
				body {
			font-family: Arial, sans-serif;
			margin: 0;
			padding: 0;
			background-color: #f4f4f4;
		}

		.email-container {
			width: 100%;
			max-width: 600px;
			margin: 0 auto;
			background-color: #ffffff;
			padding: 20px;
			border-radius: 8px;
		}

		.email-header {
			text-align: center;
			margin-bottom: 20px;
		}

		.logo {
			width: 120px;
		}

		.email-content {
			padding: 20px;
		}

		h1 {
			font-size: 24px;
			margin-bottom: 10px;
		}

		p {
			font-size: 14px;
			margin-bottom: 15px;
		}

		a {
			text-decoration: none;
			color: #007bff;
			font-weight: bold;
		}

		.manage-subscription-btn {
			background-color: #007bff;
			color: #ffffff;
			padding: 10px 15px;
			border-radius: 5px;
			display: inline-block;
			text-align: center;
			margin-bottom: 20px;
		}

		.order-details {
			background-color: #f9f9f9;
			padding: 15px;
			border-radius: 5px;
			margin-bottom: 20px;
		}

		.order-summary {
			width: 100%;
			border-collapse: collapse;
			margin-top: 10px;
		}

		.order-summary td {
			padding: 8px;
			font-size: 14px;
			text-align: left;
		}

		.order-summary tr:nth-child(even) {
			background-color: #f1f1f1;
		}

		.payment-method {
			margin-bottom: 20px;
		}

		.footer {
			font-size: 12px;
			color: #777;
			text-align: center;
		}

		.footer a {
			color: #007bff;
		}

			</style>
		</head>
		<body>
			<div class="email-container">
				<div class="email-header">
					<img src="https://app.xoailac.top/assets/images/brand/logo/logo-3.svg" alt="XCloud Logo" class="logo">
				</div>
				
				<div class="email-content">
					<h1>Bạn đã đăng ký Thành viên thành công.</h1>

					<p>Nếu bạn có bất kỳ câu hỏi nào, vui lòng liên hệ với chúng tôi thông qua telegram của chúng tôi.</p>
					
					<div class="order-details">
						<p><strong>Số đăng ký:</strong> sub_1SVBSwC6h1nxGoI3TI54ZHvc</p>
						<p><strong>Ngày đăng ký:</strong> {{date}}</p>
						<table class="order-summary">
							<tr>
								<td>Gói</td>
								<td>Giá</td>
							</tr>
							<tr>
								<td>Đăng ký thành viên</td>
								<td>{{price}}$ </td>
							</tr>
							<tr>
								<td>Thuế</td>
								<td>0$</td>
							</tr>
							<tr>
								<td>Giảm giá</td>
								<td>-9$</td>
							</tr>
							<tr>
								<td><strong>Tổng</strong></td>
								<td><strong>{{price}}$</strong></td>
							</tr>
						</table>
					</div>
					
					<div class="payment-method">
						<p><strong>Phương thức thanh toán:</strong> Paypal</p>
					</div>

					<div class="footer">
						<p>Bằng cách đăng ký, bạn cho phép chúng tôi tính phí đăng ký cho bạn như mô tả ở trên. Chi phí đăng ký sẽ được tự động tính vào phương thức thanh toán được cung cấp cho đến khi bị hủy.</p>
					</div>
				</div>
			</div>
		</body>
		</html>
		`
		data = map[string]string{
			"price":	email.Price,
			"date":		email.Date,
		}
	}
	
    for k, v := range data {
        htmlTemplate = strings.ReplaceAll(htmlTemplate, "{{"+k+"}}", v)
    }

    client := resend.NewClient(p.config.ResentApiKey)
	
    params := &resend.SendEmailRequest{
        From:    p.config.MailSender,
        To:      email.ToOfResent,
        Subject: email.Subject,
        Html:    htmlTemplate,
    }

    _, err := client.Emails.Send(params)
	if err != nil {
		
		return utils.WrapError(string(utils.ErrCodeInternal), "Failed to send request", err)
	}	

	return nil
}