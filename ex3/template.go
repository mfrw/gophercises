package ex3

var htmlTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=device-width" />
		<title>Choose Your Own Adventure</title>
	</head>
	<body>
	<section class="page">
		<h1>{{.Title}}</h1>
		{{range .Paragraphs}}
		<p>{{.}}</p>
		{{end}}

		<ul>
			{{range .Options}}
			<li><a href="/{{.Chapter}}">{{.Text}}</a></li>
			{{end}}
		</ul>
	</section>
	<style>
		body {
			font-family: helvetica, arial;
		      }
		      h1 {
			text-align:center;
			position:relative;
		      }
		      .page {
			width: 80%;
			max-width: 500px;
			margin: auto;
			margin-top: 40px;
			margin-bottom: 40px;
			padding: 80px;
			background: #FFFCF6;
			border: 1px solid #eee;
			box-shadow: 0 10px 6px -6px #777;
		      }
		      ul {
			border-top: 1px dotted #ccc;
			padding: 10px 0 0 0;
			-webkit-padding-start: 0;
		      }
		      li {
			padding-top: 10px;
		      }
		      a,
		      a:visited {
			text-decoration: none;
			color: #6295b5;
		      }
		      a:active,
		      a:hover {
			color: #7792a2;
		      }
		      p {
			text-indent: 1em;
		      }
	    </style>
	</body>
</html>
`
