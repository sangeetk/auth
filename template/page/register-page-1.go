package page

const Register1 = `
{{ define "screen" }}
<div id="login-page" class="row">

	<div class="col s12 z-depth-4 card-panel">

		<form method="POST" class="login-form" name="register" action="/user/register">
			<input id="step" name="step" type="hidden" ref="step" value="1">

			<div class="row">
				<div class="input-field col s12 center">
					<h4>Register</h4>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s6">
					<i class="mdi-social-person-outline prefix"></i>
					<input id="fname" name="fname" type="text" ref="fname" required>
					<label for="fname" class="center-align">First Name</label>
				</div>
				<div class="input-field col s6">
					<input id="lname" name="lname" type="text" ref="lname" required>
					<label for="lname" class="center-align">Last Name</label>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<i class="mdi-communication-email prefix"></i>
					<input id="email" name="email" type="email" ref="email" required>
					<label for="email" class="center-align">Email</label>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<i class="mdi-action-lock-outline prefix"></i>
					<input id="password" name="password" type="password" ref="password" required>
					<label for="password">Password</label>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<i class="mdi-action-lock-outline prefix"></i>
					<input id="password-again" type="password" ref="passwordagain" required>
					<label for="password-again">Password again</label>
				</div>
			</div>

			<div class="row">
				<div class="input-field col s12">
					<button class="btn waves-effect waves-light col s12" type="submit">Next</button>
				</div>
				<div class="input-field col s12">
					<p class="margin center medium-small sign-up">Already have an account? <a href="/user/login">Login</a></p>
				</div>
			</div>
			
		</form>

	</div>
</div>
{{ end }}
`
