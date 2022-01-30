package homeTemplate

const TPL = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<link rel="stylesheet" href="/static/css/home.css">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{.Message}}
	</body>
</html>`
