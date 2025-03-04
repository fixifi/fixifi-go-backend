package handlers

import "github.com/gofiber/fiber/v3"

func (mainHandler *MainHandler) WellComeHandler() fiber.Handler {
	return func(c fiber.Ctx) error {
		htmlContent := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Welcome to Fixifi</title>
			<style>
				@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;600&display=swap');
	
				body {
					background: linear-gradient(135deg, #ff7300, #ff3d00);
					color: #fff;
					text-align: center;
					font-family: 'Poppins', sans-serif;
					margin: 0;
					padding: 50px;
					display: flex;
					flex-direction: column;
					align-items: center;
					justify-content: center;
					height: 100vh;
				}
	
				.container {
					max-width: 600px;
					padding: 20px;
					background: rgba(0, 0, 0, 0.3);
					border-radius: 15px;
					box-shadow: 0px 10px 25px rgba(0, 0, 0, 0.4);
					animation: fadeIn 1.5s ease-in-out;
				}
	
				h1 {
					font-size: 50px;
					font-weight: 600;
					text-shadow: 2px 2px 10px rgba(255, 255, 255, 0.3);
				}
	
				h2 {
					font-size: 30px;
					font-weight: 400;
					margin-bottom: 15px;
				}
	
				h3 {
					font-size: 20px;
					font-weight: 300;
					opacity: 0.9;
				}
	
				.footer {
					margin-top: 40px;
					font-size: 18px;
					font-weight: 400;
					padding: 10px;
					background: rgba(0, 0, 0, 0.5);
					border-radius: 10px;
					box-shadow: 0px 5px 15px rgba(0, 0, 0, 0.3);
					animation: slideUp 1.5s ease-in-out;
				}
	
				@keyframes fadeIn {
					from { opacity: 0; transform: translateY(-10px); }
					to { opacity: 1; transform: translateY(0); }
				}
	
				@keyframes slideUp {
					from { opacity: 0; transform: translateY(20px); }
					to { opacity: 1; transform: translateY(0); }
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>🚩 जय श्री राम 🚩</h1>
				<h2>Welcome to Fixifi</h2>
				<h3>Qbitrom is evolving. The fusion of tech and consciousness is coming soon. 🚀</h3>
			</div>
			
			<div class="footer">
				Developed by <b>Qbitrom</b> - The Future of Innovation.
			</div>
		</body>
		</html>
		`

		return c.Type("html").SendString(htmlContent)
	}
}
