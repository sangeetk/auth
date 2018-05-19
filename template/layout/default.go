package layout

const Default = `
<!DOCTYPE html>
<html lang="en">

<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="msapplication-tap-highlight" content="no">
<meta name="description" content="">
<meta name="keywords" content="">
<title>Teaching Mission</title>

<!-- Favicons-->
<link rel="icon" href="/images/favicon/favicon-32x32.png" sizes="32x32">
<!-- Favicons-->
<link rel="apple-touch-icon-precomposed" href="/images/favicon/apple-touch-icon-152x152.png">
<!-- For iPhone -->
<meta name="msapplication-TileColor" content="#00bcd4">
<meta name="msapplication-TileImage" content="/images/favicon/mstile-144x144.png">
<!-- For Windows Phone -->

<!-- CORE CSS-->
<link href="/css/materialize.css" type="text/css" rel="stylesheet" media="screen,projection">
<link href="/css/style.css" type="text/css" rel="stylesheet" media="screen,projection">
<!-- Custome CSS--> 
<link href="/css/custom/custom.css" type="text/css" rel="stylesheet" media="screen,projection">
<link href="/css/layouts/page-center.css" type="text/css" rel="stylesheet" media="screen,projection">

<!-- INCLUDED PLUGIN CSS ON THIS PAGE -->
<link href="/js/plugins/prism/prism.css" type="text/css" rel="stylesheet" media="screen,projection">
<link href="/js/plugins/perfect-scrollbar/perfect-scrollbar.css" type="text/css" rel="stylesheet" media="screen,projection">

</head>

<body class="blue-grey">

    {{ block "screen" . }}{{ end }}

<!-- ================================================
Scripts
================================================ -->

<!-- jQuery Library -->
<script type="text/javascript" src="/js/plugins/jquery-1.11.2.min.js"></script>
<!--materialize js-->
<script type="text/javascript" src="/js/materialize.js"></script>
<!--prism-->
<script type="text/javascript" src="/js/plugins/prism/prism.js"></script>
<!--scrollbar-->
<script type="text/javascript" src="/js/plugins/perfect-scrollbar/perfect-scrollbar.min.js"></script>

<!--plugins.js - Some Specific JS codes for Plugin Settings-->
<script type="text/javascript" src="/js/plugins.js"></script>
<!--custom-script.js - Add your own theme custom JS-->
<script type="text/javascript" src="/js/custom-script.js"></script>

<!-- vue.js and axiom for ajax requests -->
<script type="text/javascript" src="/js/vue.js"></script>
<script type="text/javascript" src="/js/vee-validate.js"></script>
<!-- script type="text/javascript" src="/js/axiom.min.js"></script -->

{{ block "javascript" . }} {{ end }}

</body>
</html>
`
