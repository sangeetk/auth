package page

const LockScreen = `
{{ define "screen" }}
<div id="login-page" class="row">

	<div class="col s12 z-depth-4 card-panel">

		<form class="login-form" action="/">

			<div class="row">
				<div class="input-field col s12 center">
					<img src="images/login-logo.png" alt="" class="circle responsive-img valign profile-image-login">
					<h4 class="header">John Doe</h4>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<i class="mdi-action-lock-outline prefix"></i>
					<input id="password" type="password">
					<label for="password">Password</label>
				</div>
			</div>

			<div class="row">
				<div class="input-field col s12">
					<a href="/" class="btn waves-effect waves-light col s12">Login</a>
				</div>
			</div>

			<div class="row">
				<div class="input-field col s6 m6 l6">
					<p class="margin medium-small"><a href="/register">Register Now!</a></p>
				</div>
				<div class="input-field col s6 m6 l6">
					<p class="margin right-align medium-small"><a href="/forgot-password">Forgot password ?</a></p>
				</div>
			</div>

		</form>

	</div>
	
</div>
{{ end }}
`
