package page

const ErrorAlreadyRegistered = `
{{ define "screen" }}
<div id="login-page" class="row">

	<div class="col s12 z-depth-4 card-panel">

		<form class="login-form" action="/">

			<div class="row">
				<div class="input-field col s12 center">
					<img src="images/login-logo.png" alt="" class="circle responsive-img valign profile-image-login">
					<h4 class="header">Error !!</h4>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<p>
						Error: {{ .Err }}
					</p>
					<br>
					<br>
				</div>
			</div>

			<div class="row">
				<div class="input-field col s12">
					<a href="/forgot-password" class="btn waves-effect waves-light col s12">Forgot Password?</a>
				</div>

				<div class="input-field col s6">
					<a href="/login" class="btn waves-effect waves-light col s12">Login</a>
				</div>

				<div class="input-field col s6">
					<a href="/" class="btn waves-effect waves-light col s12">Home</a>
				</div>
			</div>



		</form>

	</div>
	
</div>
{{ end }}
`
