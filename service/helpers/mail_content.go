package helpers

import (
	"fmt"

	"github.com/DeniesKresna/bkn/models"
)

func CreateRegistrationMailContent(firstName, lastName, webURL, code string) (res string) {
	return `
		<!DOCTYPE HTML PUBLIC "-//W3C//DTD XHTML 1.0 Transitional //EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
		<html xmlns="http://www.w3.org/1999/xhtml" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office">
			<head>
			<!--[if gte mso 9]>
			<xml>
				<o:OfficeDocumentSettings>
				<o:AllowPNG/>
				<o:PixelsPerInch>96</o:PixelsPerInch>
				</o:OfficeDocumentSettings>
			</xml>
			<![endif]-->
			<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<meta name="x-apple-disable-message-reformatting">
			<!--[if !mso]><!-->
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<!--<![endif]-->
			<title></title>
			<style type="text/css">
				@media only screen and (min-width: 620px) {
				.u-row {
				width: 600px !important;
				}
				.u-row .u-col {
				vertical-align: top;
				}
				.u-row .u-col-50 {
				width: 300px !important;
				}
				.u-row .u-col-100 {
				width: 600px !important;
				}
				}
				@media (max-width: 620px) {
				.u-row-container {
				max-width: 100% !important;
				padding-left: 0px !important;
				padding-right: 0px !important;
				}
				.u-row .u-col {
				min-width: 320px !important;
				max-width: 100% !important;
				display: block !important;
				}
				.u-row {
				width: 100% !important;
				}
				.u-col {
				width: 100% !important;
				}
				.u-col > div {
				margin: 0 auto;
				}
				}
				body {
				margin: 0;
				padding: 0;
				}
				table,
				tr,
				td {
				vertical-align: top;
				border-collapse: collapse;
				}
				p {
				margin: 0;
				}
				.ie-container table,
				.mso-container table {
				table-layout: fixed;
				}
				* {
				line-height: inherit;
				}
				a[x-apple-data-detectors='true'] {
				color: inherit !important;
				text-decoration: none !important;
				}
				@media (min-width: 0px) {
				.hide-default__display-table {
				display: table !important;
				mso-hide: unset !important;
				}
				}
				@media (max-width: 480px) {
				.hide-mobile {
				max-height: 0px;
				overflow: hidden;
				display: none !important;
				}
				}
				@media (min-width: 481px) and (max-width: 768px) {
				}
				@media (min-width: 481px) {
				.hide-desktop {
				max-height: 0px;
				overflow: hidden;
				display: none !important;
				}
				}
				table, td { color: #000000; } #u_body a { color: #0000ee; text-decoration: underline; } @media (max-width: 480px) { #u_content_image_6 .v-container-padding-padding { padding: 30px 10px !important; } #u_content_text_9 .v-text-align { text-align: center !important; } #u_content_text_16 .v-text-align { text-align: center !important; } #u_content_text_13 .v-container-padding-padding { padding: 40px 10px 10px !important; } }
			</style>
			<!--[if !mso]><!-->
			<link href="https://fonts.googleapis.com/css?family=Open+Sans:400,700&display=swap" rel="stylesheet" type="text/css">
			<!--<![endif]-->
			</head>
			<body class="clean-body u_body" style="margin: 0;padding: 0;-webkit-text-size-adjust: 100%;background-color: #ffffff;color: #000000">
			<!--[if IE]>
			<div class="ie-container">
				<![endif]-->
				<!--[if mso]>
				<div class="mso-container">
				<![endif]-->
				<table id="u_body" style="border-collapse: collapse;table-layout: fixed;border-spacing: 0;mso-table-lspace: 0pt;mso-table-rspace: 0pt;vertical-align: top;min-width: 320px;Margin: 0 auto;background-color: #ffffff;width:100%" cellpadding="0" cellspacing="0">
					<tbody>
					<tr style="vertical-align: top">
						<td style="word-break: break-word;border-collapse: collapse !important;vertical-align: top">
						<!--[if (mso)|(IE)]>
						<table width="100%" cellpadding="0" cellspacing="0" border="0">
							<tr>
							<td align="center" style="background-color: #ffffff;">
								<![endif]-->
								<!--[if gte mso 9]>
								<table cellpadding="0" cellspacing="0" border="0" style="margin: 0 auto;min-width: 320px;max-width: 600px;">
								<tr>
									<td background="https://cdn.templates.unlayer.com/assets/1690713113029-bg.png" valign="top" width="100%">
									<v:rect xmlns:v="urn:schemas-microsoft-com:vml" fill="true" stroke="false" style="width: 600px;">
										<v:fill type="frame" src="https://cdn.templates.unlayer.com/assets/1690713113029-bg.png" />
										<v:textbox style="mso-fit-shape-to-text:true" inset="0,0,0,0">
										<![endif]-->
										<div class="u-row-container" style="padding: 0px;background-image: url('https://storage.googleapis.com/test-email23/image-1.png');background-repeat: no-repeat;background-position: center top;background-color: transparent">
											<div class="u-row" style="margin: 0 auto;min-width: 320px;max-width: 600px;overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;">
											<div style="border-collapse: collapse;display: table;width: 100%;height: 100%;background-color: transparent;">
												<!--[if (mso)|(IE)]>
												<table width="100%" cellpadding="0" cellspacing="0" border="0">
												<tr>
													<td style="padding: 0px;background-image: url('https://storage.googleapis.com/test-email23/image-1.png');background-repeat: no-repeat;background-position: center top;background-color: transparent;" align="center">
													<table cellpadding="0" cellspacing="0" border="0" style="width:600px;">
														<tr style="background-color: transparent;">
														<![endif]-->
														<!--[if (mso)|(IE)]>
														<td align="center" width="600" style="width: 600px;padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;" valign="top">
															<![endif]-->
															<div class="u-col u-col-100" style="max-width: 320px;min-width: 600px;display: table-cell;vertical-align: top;">
															<div style="height: 100%;width: 100% !important;">
																<!--[if (!mso)&(!IE)]><!-->
																<div style="box-sizing: border-box; height: 100%; padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;">
																<!--<![endif]-->
																<table style="font-family:'Open Sans',sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
																	<tbody>
																	<tr>
																		<td class="v-container-padding-padding" style="overflow-wrap:break-word;word-break:break-word;padding:0px;font-family:'Open Sans',sans-serif;" align="left">
																		<table width="100%" cellpadding="0" cellspacing="0" border="0">
																			<tr>
																			<td class="v-text-align" style="padding-right: 0px;padding-left: 0px;" align="center">
																				<img align="center" border="0" src="https://storage.googleapis.com/test-email23/image-2.jpeg" alt="image" title="image" style="outline: none;text-decoration: none;-ms-interpolation-mode: bicubic;clear: both;display: inline-block !important;border: none;height: auto;float: none;width: 100%;max-width: 600px;" width="600"/>
																			</td>
																			</tr>
																		</table>
																		</td>
																	</tr>
																	</tbody>
																</table>
																<!--[if (!mso)&(!IE)]><!-->
																</div>
																<!--<![endif]-->
															</div>
															</div>
															<!--[if (mso)|(IE)]>
														</td>
														<![endif]-->
														<!--[if (mso)|(IE)]>
														</tr>
													</table>
													</td>
												</tr>
												</table>
												<![endif]-->
											</div>
											</div>
										</div>
										<!--[if gte mso 9]>
										</v:textbox>
									</v:rect>
									</td>
								</tr>
								</table>
								<![endif]-->
								<div class="u-row-container" style="padding: 20px 0px 0px;background-color: transparent">
								<div class="u-row" style="margin: 0 auto;min-width: 320px;max-width: 600px;overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;">
									<div style="border-collapse: collapse;display: table;width: 100%;height: 100%;background-color: transparent;">
									<!--[if (mso)|(IE)]>
									<table width="100%" cellpadding="0" cellspacing="0" border="0">
										<tr>
										<td style="padding: 20px 0px 0px;background-color: transparent;" align="center">
											<table cellpadding="0" cellspacing="0" border="0" style="width:600px;">
											<tr style="background-color: transparent;">
												<![endif]-->
												<!--[if (mso)|(IE)]>
												<td align="center" width="298" style="width: 298px;padding: 0px;border-top: 1px solid transparent;border-left: 1px solid transparent;border-right: 1px solid transparent;border-bottom: 1px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;" valign="top">
												<![endif]-->
												<div class="u-col u-col-50" style="max-width: 320px;min-width: 300px;display: table-cell;vertical-align: top;">
													<div style="height: 100%;width: 100% !important;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">
													<!--[if (!mso)&(!IE)]><!-->
													<div style="box-sizing: border-box; height: 100%; padding: 0px;border-top: 1px solid transparent;border-left: 1px solid transparent;border-right: 1px solid transparent;border-bottom: 1px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">
														<!--<![endif]-->
														<table id="u_content_image_6" style="font-family:'Open Sans',sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
														<tbody>
															<tr>
															<td class="v-container-padding-padding" style="overflow-wrap:break-word;word-break:break-word;padding:45px 10px 10px;font-family:'Open Sans',sans-serif;" align="left">
																<table width="100%" cellpadding="0" cellspacing="0" border="0">
																<tr>
																	<td class="v-text-align" style="padding-right: 0px;padding-left: 0px;" align="center">
																	<img align="center" border="0" src="https://storage.googleapis.com/test-email23/image-3.png" alt="image" title="image" style="outline: none;text-decoration: none;-ms-interpolation-mode: bicubic;clear: both;display: inline-block !important;border: none;height: auto;float: none;width: 70%;max-width: 196px;" width="196"/>
																	</td>
																</tr>
																</table>
															</td>
															</tr>
														</tbody>
														</table>
														<table class="hide-mobile" style="font-family:'Open Sans',sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
														<tbody>
															<tr>
															<td class="v-container-padding-padding" style="overflow-wrap:break-word;word-break:break-word;padding:10px;font-family:'Open Sans',sans-serif;" align="left">
																<div class="v-text-align" style="font-size: 14px; line-height: 140%; text-align: center; word-wrap: break-word;">
																<p style="line-height: 140%;"><span style="color: #414141; line-height: 19.6px;"><strong>Jobhun -</strong></span></p>
																<p style="line-height: 140%;"><span style="color: #414141; line-height: 19.6px;"><strong>PT Jobhun Membangun Indonesia</strong></span></p>
																</div>
															</td>
															</tr>
														</tbody>
														</table>
														<table class="hide-mobile" style="font-family:'Open Sans',sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
														<tbody>
															<tr>
															<td class="v-container-padding-padding" style="overflow-wrap:break-word;word-break:break-word;padding:10px;font-family:'Open Sans',sans-serif;" align="left">
																<div class="v-text-align" style="font-size: 14px; line-height: 140%; text-align: left; word-wrap: break-word;">
																<p style="line-height: 140%;"><span style="color: #414141; line-height: 19.6px;">Gedung Inkubator Bisnis Universitas Airlangga, Jl. Dharmawangsa No. 33B, Airlangga, Gubeng, Surabaya, Jawa Timur, Indonesia 60286</span></p>
																<p style="line-height: 140%;"> </p>
																<p style="line-height: 140%;"><span style="color: #414141; line-height: 19.6px;">HQ: Maspion Square, M.IT 3 Blok F.08, Jl. Ahmad Yani 73, Margorejo, Wonocolo, Surabaya, Jawa Timur, Indonesia 60238</span></p>
																<p style="line-height: 140%;"> </p>
																<p style="line-height: 140%;"><span style="color: #414141; line-height: 19.6px;">E-mail: info@jobhun.id; partnership@jobhun.id | Phone: 081252982369</span></p>
																</div>
															</td>
															</tr>
														</tbody>
														</table>
														<!--[if (!mso)&(!IE)]><!-->
													</div>
													<!--<![endif]-->
													</div>
												</div>
												<!--[if (mso)|(IE)]>
												</td>
												<![endif]-->
												<!--[if (mso)|(IE)]>
												<td align="center" width="300" style="width: 300px;padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;" valign="top">
												<![endif]-->
												<div class="u-col u-col-50" style="max-width: 320px;min-width: 300px;display: table-cell;vertical-align: top;">
													<div style="height: 100%;width: 100% !important;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">
													<!--[if (!mso)&(!IE)]><!-->
													<div style="box-sizing: border-box; height: 100%; padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">
														<!--<![endif]-->
														<table id="u_content_text_9" style="font-family:'Open Sans',sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
														<tbody>
															<tr>
															<td class="v-container-padding-padding" style="overflow-wrap:break-word;word-break:break-word;padding:5px 5px 0px;font-family:'Open Sans',sans-serif;" align="left">
																<div class="v-text-align" style="font-size: 14px; color: #ebebeb; line-height: 140%; text-align: left; word-wrap: break-word;">
																<p style="line-height: 140%;"><span style="color: #000000; line-height: 19.6px;">Hai <strong>` + firstName + ` ` + lastName + `</strong>,</span></p>
																<p style="line-height: 140%;"><span style="color: #000000; line-height: 19.6px;">Selamat datang di Jobhun! Kami senang menyambutmu menjadi bagian dari ribuan pengguna Jobhun. Kamu bisa menemukan ratusan expert di Jobhun untuk berbagai kebutuhan pengembangan karier seperti konsultasi, pelatihan, undang expert, dan rekrut expert. Segera verifikasi akunmu dengan menekan tombol verifikasi untuk dapat menikmati akses tanpa batas ke semua fitur Jobhun.</span></p>
																<p style="line-height: 140%;"> </p>
																</div>
															</td>
															</tr>
														</tbody>
														</table>
														<table style="font-family:'Open Sans',sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
														<tbody>
															<tr>
															<td class="v-container-padding-padding" style="overflow-wrap:break-word;word-break:break-word;padding:10px;font-family:'Open Sans',sans-serif;" align="left">
																<!--[if mso]>
																<style>.v-button {background: transparent !important;}</style>
																<![endif]-->
																<div class="v-text-align" align="center">
																<!--[if mso]>
																<v:roundrect xmlns:v="urn:schemas-microsoft-com:vml" xmlns:w="urn:schemas-microsoft-com:office:word" href="" style="height:37px; v-text-anchor:middle; width:135px;" arcsize="11%"  stroke="f" fillcolor="#48b391">
																	<w:anchorlock/>
																	<center style="color:#FFFFFF;">
																	<![endif]-->
																	<a href="` + webURL + `/signin?veriftoken=` + code + `" target="_blank" class="v-button" style="box-sizing: border-box;display: inline-block;text-decoration: none;-webkit-text-size-adjust: none;text-align: center;color: #FFFFFF; background-color: #48b391; border-radius: 4px;-webkit-border-radius: 4px; -moz-border-radius: 4px; width:auto; max-width:100%; overflow-wrap: break-word; word-break: break-word; word-wrap:break-word; mso-border-alt: none;font-size: 14px;">
																	<span style="display:block;padding:10px 20px;line-height:120%;"><span style="line-height: 16.8px;">Verifikasi akun</span></span>
																	</a>
																	<!--[if mso]>
																	</center>
																</v:roundrect>
																<![endif]-->
																</div>
															</td>
															</tr>
														</tbody>
														</table>
														<table id="u_content_text_16" style="font-family:'Open Sans',sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
														<tbody>
															<tr>
															<td class="v-container-padding-padding" style="overflow-wrap:break-word;word-break:break-word;padding:10px;font-family:'Open Sans',sans-serif;" align="left">
																<div class="v-text-align" style="font-size: 14px; color: #ebebeb; line-height: 140%; text-align: left; word-wrap: break-word;">
																<p style="line-height: 140%;"><span style="color: #000000; line-height: 19.6px;">Tombol verifikasi ini hanya berlaku selama 1 x 24 Jam sejak pertama kali email ini diterima. Mohon jangan sebarkan kode verifikasi ini ke siapa pun, termasuk ke pihak yang mengatasnamakan Jobhun.</span></p>
																<p style="line-height: 140%;"><span style="color: #000000; line-height: 19.6px;">Apabila ada kendala, saran, atau kritik, hubungi Jobhun melalui <strong><span style="text-decoration: underline; line-height: 19.6px;"><span style="color: #48b391; line-height: 19.6px; text-decoration: underline;">info@jobhun.id</span></span> </strong>atau WhatsApp 08125982369.</span></p>
																<p style="line-height: 140%;"><span style="color: #000000; line-height: 19.6px;">Selamat berproses bersama!</span></p>
																<p style="line-height: 140%;"> </p>
																<p style="line-height: 140%;"><span style="color: #000000; line-height: 19.6px;">Salam,</span></p>
																<p style="line-height: 140%;"><span style="color: #000000; line-height: 19.6px;">Jobhun Squad</span></p>
																<p style="line-height: 140%;"><br /><br /></p>
																</div>
															</td>
															</tr>
														</tbody>
														</table>
														<!--[if (!mso)&(!IE)]><!-->
													</div>
													<!--<![endif]-->
													</div>
												</div>
												<!--[if (mso)|(IE)]>
												</td>
												<![endif]-->
												<!--[if (mso)|(IE)]>
											</tr>
											</table>
										</td>
										</tr>
									</table>
									<![endif]-->
									</div>
								</div>
								</div>
								<div class="u-row-container" style="padding: 0px;background-color: transparent">
								<div class="u-row" style="margin: 0 auto;min-width: 320px;max-width: 600px;overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;">
									<div style="border-collapse: collapse;display: table;width: 100%;height: 100%;background-color: transparent;">
									<!--[if (mso)|(IE)]>
									<table width="100%" cellpadding="0" cellspacing="0" border="0">
										<tr>
										<td style="padding: 0px;background-color: transparent;" align="center">
											<table cellpadding="0" cellspacing="0" border="0" style="width:600px;">
											<tr style="background-color: transparent;">
												<![endif]-->
												<!--[if (mso)|(IE)]>
												<td align="center" width="600" style="width: 600px;padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;" valign="top">
												<![endif]-->
												<div class="u-col u-col-100" style="max-width: 320px;min-width: 600px;display: table-cell;vertical-align: top;">
													<div style="height: 100%;width: 100% !important;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">
													<!--[if (!mso)&(!IE)]><!-->
													<div style="box-sizing: border-box; height: 100%; padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">
														<!--<![endif]-->
														<!--[if !mso]><!-->
														<table id="u_content_text_13" class="hide-default__display-table hide-desktop" style="display: none;mso-hide: all;font-family:'Open Sans',sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
														<tbody>
															<tr>
															<td class="v-container-padding-padding" style="overflow-wrap:break-word;word-break:break-word;padding:40px 80px 10px;font-family:'Open Sans',sans-serif;" align="left">
																<div class="v-text-align" style="font-size: 14px; line-height: 160%; text-align: center; word-wrap: break-word;">
																<p style="line-height: 160%;"><span style="color: #414141; line-height: 22.4px;"><strong>Jobhun -</strong></span></p>
																<p style="line-height: 160%;"><span style="color: #414141; line-height: 22.4px;"><strong>PT Jobhun Membangun Indonesia</strong></span></p>
																</div>
															</td>
															</tr>
														</tbody>
														</table>
														<!--<![endif]--><!--[if !mso]><!-->
														<table class="hide-default__display-table hide-desktop" style="display: none;mso-hide: all;font-family:'Open Sans',sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
														<tbody>
															<tr>
															<td class="v-container-padding-padding" style="overflow-wrap:break-word;word-break:break-word;padding:10px 10px 40px;font-family:'Open Sans',sans-serif;" align="left">
																<div class="v-text-align" style="font-size: 14px; line-height: 160%; text-align: center; word-wrap: break-word;">
																<p style="line-height: 160%;"><span style="color: #414141; line-height: 22.4px;">Gedung Inkubator Bisnis Universitas Airlangga, Jl. Dharmawangsa No. 33B, Airlangga, Gubeng, Surabaya, Jawa Timur, Indonesia 60286</span></p>
																<p style="line-height: 160%;"> </p>
																<p style="line-height: 160%;"><span style="color: #414141; line-height: 22.4px;">HQ: Maspion Square, M.IT 3 Blok F.08, Jl. Ahmad Yani 73, Margorejo, Wonocolo, Surabaya, Jawa Timur, Indonesia 60238</span></p>
																<p style="line-height: 160%;"> </p>
																<p style="line-height: 160%;"><span style="color: #414141; line-height: 22.4px;">E-mail: info@jobhun.id; partnership@jobhun.id | Phone: 081252982369</span></p>
																</div>
															</td>
															</tr>
														</tbody>
														</table>
														<!--<![endif]-->
														<!--[if (!mso)&(!IE)]><!-->
													</div>
													<!--<![endif]-->
													</div>
												</div>
												<!--[if (mso)|(IE)]>
												</td>
												<![endif]-->
												<!--[if (mso)|(IE)]>
											</tr>
											</table>
										</td>
										</tr>
									</table>
									<![endif]-->
									</div>
								</div>
								</div>
								<!--[if (mso)|(IE)]>
							</td>
							</tr>
						</table>
						<![endif]-->
						</td>
					</tr>
					</tbody>
				</table>
				<!--[if mso]>
				</div>
				<![endif]-->
				<!--[if IE]>
			</div>
			<![endif]-->
			</body>
		</html>
	`
}

func CreateCoursePaymentCreationMailContentToAdmin(
	course models.Course, payment models.Payment, user models.User,
) string {
	var payAmount int64 = 0
	if payment.PayAmount != nil {
		payAmount = *payment.PayAmount
	}

	payMethod := ""
	if payment.PayMethod != nil {
		payMethod = *payment.PayMethod
	}

	payChannel := ""
	if payment.PayChannel != nil {
		payChannel = *payment.PayChannel
	}

	style := `style="line-height: 140%;"`
	contentData := fmt.Sprintf(`
	<p %s">Berikut untuk data lengkapnya:</p>
	<p %s"> </p>
	<p %s">Nama Pelatihan : %s</p>
	<p %s">Nama : %s %s</p>
	<p %s">Email : %s</p>
	<p %s">Nomor WhatsApp : %s</p>
	<p %s">Periode Kelas : %s - %s</p>
	<p %s">Jumlah Pembayaran : %d</p>
	<p %s">Dibayar pada : %s</p>
	<p %s">Kode Pembayaran Vendor : %s</p>
	<p %s">Metode Pembayaran : %s</p>
	<p %s">Saluran Pembayaran : %s</p>
	`,
		style, style, style, course.Name, style, user.FirstName, user.LastName, style, user.Email, style,
		user.Phone, style, course.StartTime, course.EndTime, style, payAmount, style, payment.PayAt, style,
		payment.VendorID, style, payMethod, style, payChannel)

	return `
	<!DOCTYPE HTML PUBLIC "-//W3C//DTD XHTML 1.0 Transitional //EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
	<html xmlns="http://www.w3.org/1999/xhtml" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office">
		<head>
			<!--[if gte mso 9]>
			<xml>
				<o:OfficeDocumentSettings>
					<o:AllowPNG/>
					<o:PixelsPerInch>96</o:PixelsPerInch>
				</o:OfficeDocumentSettings>
			</xml>
			<![endif]-->
			<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<meta name="x-apple-disable-message-reformatting">
			<!--[if !mso]><!-->
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<!--<![endif]-->
			<title></title>
			<style type="text/css">
				@media only screen and (min-width: 620px) {
				.u-row {
				width: 600px !important;
				}
				.u-row .u-col {
				vertical-align: top;
				}
				.u-row .u-col-50 {
				width: 300px !important;
				}
				.u-row .u-col-100 {
				width: 600px !important;
				}
				}
				@media (max-width: 620px) {
				.u-row-container {
				max-width: 100% !important;
				padding-left: 0px !important;
				padding-right: 0px !important;
				}
				.u-row .u-col {
				min-width: 320px !important;
				max-width: 100% !important;
				display: block !important;
				}
				.u-row {
				width: 100% !important;
				}
				.u-col {
				width: 100% !important;
				}
				.u-col > div {
				margin: 0 auto;
				}
				}
				body {
				margin: 0;
				padding: 0;
				}
				table,
				tr,
				td {
				vertical-align: top;
				border-collapse: collapse;
				}
				p {
				margin: 0;
				}
				.ie-container table,
				.mso-container table {
				table-layout: fixed;
				}
				* {
				line-height: inherit;
				}
				a[x-apple-data-detectors='true'] {
				color: inherit !important;
				text-decoration: none !important;
				}
				@media (min-width: 481px) and (max-width: 768px) {
				}
				table, td { color: #000000; } 
			</style>
		</head>
		<body class="clean-body u_body" style="margin: 0;padding: 0;-webkit-text-size-adjust: 100%;background-color: #e7e7e7;color: #000000">
			<!--[if IE]>
			<div class="ie-container">
				<![endif]-->
				<!--[if mso]>
				<div class="mso-container">
					<![endif]-->
					<table style="border-collapse: collapse;table-layout: fixed;border-spacing: 0;mso-table-lspace: 0pt;mso-table-rspace: 0pt;vertical-align: top;min-width: 320px;Margin: 0 auto;background-color: #e7e7e7;width:100%" cellpadding="0" cellspacing="0">
						<tbody>
							<tr style="vertical-align: top">
								<td style="word-break: break-word;border-collapse: collapse !important;vertical-align: top">
									<!--[if (mso)|(IE)]>
									<table width="100%" cellpadding="0" cellspacing="0" border="0">
										<tr>
											<td align="center" style="background-color: #e7e7e7;">
												<![endif]-->
												<div class="u-row-container" style="padding: 0px;background-color: transparent">
													<div class="u-row" style="margin: 0 auto;min-width: 320px;max-width: 600px;overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;">
														<div style="border-collapse: collapse;display: table;width: 100%;height: 100%;background-color: transparent;">
															<!--[if (mso)|(IE)]>
															<table width="100%" cellpadding="0" cellspacing="0" border="0">
																<tr>
																	<td style="padding: 0px;background-color: transparent;" align="center">
																		<table cellpadding="0" cellspacing="0" border="0" style="width:600px;">
																			<tr style="background-color: transparent;">
																				<![endif]-->
																				<!--[if (mso)|(IE)]>
																				<td align="center" width="600" style="width: 600px;padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;" valign="top">
																					<![endif]-->
																					<div class="u-col u-col-100" style="max-width: 320px;min-width: 600px;display: table-cell;vertical-align: top;">
																						<div style="height: 100%;width: 100% !important;">
																							<!--[if (!mso)&(!IE)]><!-->
																							<div style="box-sizing: border-box; height: 100%; padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;">
																								<!--<![endif]-->
																								<table style="font-family:arial,helvetica,sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
																									<tbody>
																										<tr>
																											<td style="overflow-wrap:break-word;word-break:break-word;padding:0px;font-family:arial,helvetica,sans-serif;" align="left">
																												<table width="100%" cellpadding="0" cellspacing="0" border="0">
																													<tr>
																														<td style="padding-right: 0px;padding-left: 0px;" align="center">
																															<img align="center" border="0" src="images/image-2.jpeg" alt="image" title="image" style="outline: none;text-decoration: none;-ms-interpolation-mode: bicubic;clear: both;display: inline-block !important;border: none;height: auto;float: none;width: 100%;max-width: 600px;" width="600"/>
																														</td>
																													</tr>
																												</table>
																											</td>
																										</tr>
																									</tbody>
																								</table>
																								<!--[if (!mso)&(!IE)]><!-->
																							</div>
																							<!--<![endif]-->
																						</div>
																					</div>
																					<!--[if (mso)|(IE)]>
																				</td>
																				<![endif]-->
																				<!--[if (mso)|(IE)]>
																			</tr>
																		</table>
																	</td>
																</tr>
															</table>
															<![endif]-->
														</div>
													</div>
												</div>
												<div class="u-row-container" style="padding: 0px;background-color: transparent">
													<div class="u-row" style="margin: 0 auto;min-width: 320px;max-width: 600px;overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;">
														<div style="border-collapse: collapse;display: table;width: 100%;height: 100%;background-color: transparent;">
															<!--[if (mso)|(IE)]>
															<table width="100%" cellpadding="0" cellspacing="0" border="0">
																<tr>
																	<td style="padding: 0px;background-color: transparent;" align="center">
																		<table cellpadding="0" cellspacing="0" border="0" style="width:600px;">
																			<tr style="background-color: transparent;">
																				<![endif]-->
																				<!--[if (mso)|(IE)]>
																				<td align="center" width="300" style="background-color: #ffffff;width: 300px;padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;" valign="top">
																					<![endif]-->
																					<div class="u-col u-col-50" style="max-width: 320px;min-width: 300px;display: table-cell;vertical-align: top;">
																						<div style="background-color: #ffffff;height: 100%;width: 100% !important;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">
																							<!--[if (!mso)&(!IE)]><!-->
																							<div style="box-sizing: border-box; height: 100%; padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">
																								<!--<![endif]-->
																								<table style="font-family:arial,helvetica,sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
																									<tbody>
																										<tr>
																											<td style="overflow-wrap:break-word;word-break:break-word;padding:40px 10px 10px;font-family:arial,helvetica,sans-serif;" align="left">
																												<table width="100%" cellpadding="0" cellspacing="0" border="0">
																													<tr>
																														<td style="padding-right: 0px;padding-left: 0px;" align="center">
																															<img align="center" border="0" src="images/image-1.png" alt="image" title="image" style="outline: none;text-decoration: none;-ms-interpolation-mode: bicubic;clear: both;display: inline-block !important;border: none;height: auto;float: none;width: 52%;max-width: 145.6px;" width="145.6"/>
																														</td>
																													</tr>
																												</table>
																											</td>
																										</tr>
																									</tbody>
																								</table>
																								<!--[if (!mso)&(!IE)]><!-->
																							</div>
																							<!--<![endif]-->
																						</div>
																					</div>
																					<!--[if (mso)|(IE)]>
																				</td>
																				<![endif]-->
																				<!--[if (mso)|(IE)]>
																				<td align="center" width="300" style="background-color: #ffffff;width: 300px;padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;" valign="top">
																					<![endif]-->
																					<div class="u-col u-col-50" style="max-width: 320px;min-width: 300px;display: table-cell;vertical-align: top;">
																						<div style="background-color: #ffffff;height: 100%;width: 100% !important;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">
																							<!--[if (!mso)&(!IE)]><!-->
																							<div style="box-sizing: border-box; height: 100%; padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">
																								<!--<![endif]-->
																								<table style="font-family:arial,helvetica,sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
																									<tbody>
																										<tr>
																											<td style="overflow-wrap:break-word;word-break:break-word;padding:20px 10px 10px;font-family:arial,helvetica,sans-serif;" align="left">
																												<div style="font-size: 17px; line-height: 140%; text-align: left; word-wrap: break-word;">
																													<p style="line-height: 140%;"><strong>Yey, ada pendaftar baru! </strong></p>
																												</div>
																											</td>
																										</tr>
																									</tbody>
																								</table>
																								<table style="font-family:arial,helvetica,sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
																									<tbody>
																										<tr>
																											<td style="overflow-wrap:break-word;word-break:break-word;padding:0px 10px 20px;font-family:arial,helvetica,sans-serif;" align="left">
																												<div style="font-size: 14px; line-height: 140%; text-align: left; word-wrap: break-word;">
																													` + contentData + `
																												</div>
																											</td>
																										</tr>
																									</tbody>
																								</table>
																								<!--[if (!mso)&(!IE)]><!-->
																							</div>
																							<!--<![endif]-->
																						</div>
																					</div>
																					<!--[if (mso)|(IE)]>
																				</td>
																				<![endif]-->
																				<!--[if (mso)|(IE)]>
																			</tr>
																		</table>
																	</td>
																</tr>
															</table>
															<![endif]-->
														</div>
													</div>
												</div>
												<!--[if (mso)|(IE)]>
											</td>
										</tr>
									</table>
									<![endif]-->
								</td>
							</tr>
						</tbody>
					</table>
					<!--[if mso]>
				</div>
				<![endif]-->
				<!--[if IE]>
			</div>
			<![endif]-->
		</body>
	</html>
	`
}

func PasswordConfirmationMailContent(name, password string) (res string) {
	return `
		<!DOCTYPE HTML PUBLIC "-//W3C//DTD XHTML 1.0 Transitional //EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
		<html xmlns="http://www.w3.org/1999/xhtml" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office">
		<head>
			<!--[if gte mso 9]>
			<xml>
				<o:OfficeDocumentSettings>
					<o:AllowPNG/>
					<o:PixelsPerInch>96</o:PixelsPerInch>
				</o:OfficeDocumentSettings>
			</xml>
			<![endif]-->
			<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<meta name="x-apple-disable-message-reformatting">
			<!--[if !mso]><!-->
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<!--<![endif]-->
			<title></title>
			<style type="text/css">
				@media only screen and (min-width: 620px) {
				.u-row {
				width: 600px !important;
				}
				.u-row .u-col {
				vertical-align: top;
				}
				.u-row .u-col-50 {
				width: 300px !important;
				}
				.u-row .u-col-100 {
				width: 600px !important;
				}
				}
				@media (max-width: 620px) {
				.u-row-container {
				max-width: 100% !important;
				padding-left: 0px !important;
				padding-right: 0px !important;
				}
				.u-row .u-col {
				min-width: 320px !important;
				max-width: 100% !important;
				display: block !important;
				}
				.u-row {
				width: 100% !important;
				}
				.u-col {
				width: 100% !important;
				}
				.u-col > div {
				margin: 0 auto;
				}
				}
				body {
				margin: 0;
				padding: 0;
				}
				table,
				tr,
				td {
				vertical-align: top;
				border-collapse: collapse;
				}
				p {
				margin: 0;
				}
				.ie-container table,
				.mso-container table {
				table-layout: fixed;
				}
				* {
				line-height: inherit;
				}
				a[x-apple-data-detectors='true'] {
				color: inherit !important;
				text-decoration: none !important;
				}
				@media (min-width: 481px) and (max-width: 768px) {
				}
				table, td { color: #000000; } 
			</style>
		</head>
		<body class="clean-body u_body" style="margin: 0;padding: 0;-webkit-text-size-adjust: 100%;background-color: #e7e7e7;color: #000000">
			<!--[if IE]>
			<div class="ie-container">
				<![endif]-->
				<!--[if mso]>
				<div class="mso-container">
					<![endif]-->
					<table style="border-collapse: collapse;table-layout: fixed;border-spacing: 0;mso-table-lspace: 0pt;mso-table-rspace: 0pt;vertical-align: top;min-width: 320px;Margin: 0 auto;background-color: #e7e7e7;width:100%" cellpadding="0" cellspacing="0">
						<tbody>
							<tr style="vertical-align: top">
								<td style="word-break: break-word;border-collapse: collapse !important;vertical-align: top">
									<!--[if (mso)|(IE)]>
									<table width="100%" cellpadding="0" cellspacing="0" border="0">
										<tr>
											<td align="center" style="background-color: #e7e7e7;">
												<![endif]-->
												<div class="u-row-container" style="padding: 0px;background-color: transparent">
													<div class="u-row" style="margin: 0 auto;min-width: 320px;max-width: 600px;overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;">
														<div style="border-collapse: collapse;display: table;width: 100%;height: 100%;background-color: transparent;">
															<!--[if (mso)|(IE)]>
															<table width="100%" cellpadding="0" cellspacing="0" border="0">
																<tr>
																	<td style="padding: 0px;background-color: transparent;" align="center">
																		<table cellpadding="0" cellspacing="0" border="0" style="width:600px;">
																			<tr style="background-color: transparent;">
																				<![endif]-->
																				<!--[if (mso)|(IE)]>
																				<td align="center" width="600" style="width: 600px;padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;" valign="top">
																					<![endif]-->
																					<div class="u-col u-col-100" style="max-width: 320px;min-width: 600px;display: table-cell;vertical-align: top;">
																						<div style="height: 100%;width: 100% !important;">
																							<!--[if (!mso)&(!IE)]><!-->
																							<div style="box-sizing: border-box; height: 100%; padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;">
																								<!--<![endif]-->
																								<table style="font-family:arial,helvetica,sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
																									<tbody>
																										<tr>
																											<td style="overflow-wrap:break-word;word-break:break-word;padding:0px;font-family:arial,helvetica,sans-serif;" align="left">
																												<table width="100%" cellpadding="0" cellspacing="0" border="0">
																													<tr>
																														<td style="padding-right: 0px;padding-left: 0px;" align="center">
																															<img align="center" border="0" src="https://storage.googleapis.com/test-email23/image-2.jpeg" alt="image" title="image" style="outline: none;text-decoration: none;-ms-interpolation-mode: bicubic;clear: both;display: inline-block !important;border: none;height: auto;float: none;width: 100%;max-width: 600px;" width="600"/>
																														</td>
																													</tr>
																												</table>
																											</td>
																										</tr>
																									</tbody>
																								</table>
																								<!--[if (!mso)&(!IE)]><!-->
																							</div>
																							<!--<![endif]-->
																						</div>
																					</div>
																					<!--[if (mso)|(IE)]>
																				</td>
																				<![endif]-->
																				<!--[if (mso)|(IE)]>
																			</tr>
																		</table>
																	</td>
																</tr>
															</table>
															<![endif]-->
														</div>
													</div>
												</div>
												<div class="u-row-container" style="padding: 0px;background-color: transparent">
													<div class="u-row" style="margin: 0 auto;min-width: 320px;max-width: 600px;overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;">
														<div style="border-collapse: collapse;display: table;width: 100%;height: 100%;background-color: transparent;">
															<!--[if (mso)|(IE)]>
															<table width="100%" cellpadding="0" cellspacing="0" border="0">
																<tr>
																	<td style="padding: 0px;background-color: transparent;" align="center">
																		<table cellpadding="0" cellspacing="0" border="0" style="width:600px;">
																			<tr style="background-color: transparent;">
																				<![endif]-->
																				<!--[if (mso)|(IE)]>
																				<td align="center" width="300" style="background-color: #ffffff;width: 300px;padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;" valign="top">
																					<![endif]-->
																					<div class="u-col u-col-50" style="max-width: 320px;min-width: 300px;display: table-cell;vertical-align: top;">
																						<div style="background-color: #ffffff;height: 100%;width: 100% !important;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">
																							<!--[if (!mso)&(!IE)]><!-->
																							<div style="box-sizing: border-box; height: 100%; padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">
																								<!--<![endif]-->
																								<table style="font-family:arial,helvetica,sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
																									<tbody>
																										<tr>
																											<td style="overflow-wrap:break-word;word-break:break-word;padding:40px 10px 10px;font-family:arial,helvetica,sans-serif;" align="left">
																												<table width="100%" cellpadding="0" cellspacing="0" border="0">
																													<tr>
																														<td style="padding-right: 0px;padding-left: 0px;" align="center">
																															<img align="center" border="0" src="https://storage.googleapis.com/test-email23/image-3.png" alt="image" title="image" style="outline: none;text-decoration: none;-ms-interpolation-mode: bicubic;clear: both;display: inline-block !important;border: none;height: auto;float: none;width: 52%;max-width: 145.6px;" width="145.6"/>
																														</td>
																													</tr>
																												</table>
																											</td>
																										</tr>
																									</tbody>
																								</table>
																								<!--[if (!mso)&(!IE)]><!-->
																							</div>
																							<!--<![endif]-->
																						</div>
																					</div>
																					<!--[if (mso)|(IE)]>
																				</td>
																				<![endif]-->
																				<!--[if (mso)|(IE)]>
																				<td align="center" width="300" style="background-color: #ffffff;width: 300px;padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;" valign="top">
																					<![endif]-->
																					<div class="u-col u-col-50" style="max-width: 320px;min-width: 300px;display: table-cell;vertical-align: top;">
																						<div style="background-color: #ffffff;height: 100%;width: 100% !important;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">
																							<!--[if (!mso)&(!IE)]><!-->
																							<div style="box-sizing: border-box; height: 100%; padding: 0px;border-top: 0px solid transparent;border-left: 0px solid transparent;border-right: 0px solid transparent;border-bottom: 0px solid transparent;border-radius: 0px;-webkit-border-radius: 0px; -moz-border-radius: 0px;">
																								<!--<![endif]-->
																								<table style="font-family:arial,helvetica,sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
																									<tbody>
																										<tr>
																											<td style="overflow-wrap:break-word;word-break:break-word;padding:20px 10px 10px;font-family:arial,helvetica,sans-serif;" align="left">
																												<div style="font-size: 17px; line-height: 140%; text-align: left; word-wrap: break-word;">
																													<p style="line-height: 140%;"><strong>Halo, ` + name + `!</strong></p>
																												</div>
																											</td>
																										</tr>
																									</tbody>
																								</table>
																								<table style="font-family:arial,helvetica,sans-serif;" role="presentation" cellpadding="0" cellspacing="0" width="100%" border="0">
																									<tbody>
																										<tr>
																											<td style="overflow-wrap:break-word;word-break:break-word;padding:0px 10px 20px;font-family:arial,helvetica,sans-serif;" align="left">
																												<div style="font-size: 14px; line-height: 140%; text-align: left; word-wrap: break-word;">
																													<p style="line-height: 140%;">Terima kasih sudah mendaftar sebagai expert di Jobhun. Akunmu sebagai expert telah berhasil dibuat. Gunakan password ini untuk masuk ke website Jobhun sebagai expert. Pasword : <span style="color: red;">` + password + `</span> . Jangan lupa untuk mengganti <em>password</em>-mu setelah masuk ke akun, ya! </p>
																													<p style="line-height: 140%;"> </p>
																													<p style="line-height: 140%;">Best regards,</p>
																													<p style="line-height: 140%;">Jobhun</p>
																												</div>
																											</td>
																										</tr>
																									</tbody>
																								</table>
																								<!--[if (!mso)&(!IE)]><!-->
																							</div>
																							<!--<![endif]-->
																						</div>
																					</div>
																					<!--[if (mso)|(IE)]>
																				</td>
																				<![endif]-->
																				<!--[if (mso)|(IE)]>
																			</tr>
																		</table>
																	</td>
																</tr>
															</table>
															<![endif]-->
														</div>
													</div>
												</div>
												<!--[if (mso)|(IE)]>
											</td>
										</tr>
									</table>
									<![endif]-->
								</td>
							</tr>
						</tbody>
					</table>
					<!--[if mso]>
				</div>
				<![endif]-->
				<!--[if IE]>
			</div>
			<![endif]-->
		</body>
		</html>
	`
}
