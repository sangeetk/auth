package page

const ThankYou = `
{{ define "screen" }}
<div id="login-page" class="row">

	<div class="col s12 z-depth-4 card-panel">

		<form class="login-form" action="/">

			<div class="row">
				<div class="input-field col s12 center">
					<img src="images/login-logo.png" alt="" class="circle responsive-img valign profile-image-login">
					<h4 class="header">Thank You !!</h4>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<p>
						Further instruction to activate your account has been emailed to you.
					</p>
					<br>
					<br>
				</div>
			</div>

			<div class="row">
				<div class="input-field col s6">
					<a href="/user/login" class="btn waves-effect waves-light col s12">Login</a>
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
