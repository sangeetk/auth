package page

const Register2 = `
{{ define "screen" }}
<div id="login-page" class="row">

	<div class="col s12 z-depth-4 card-panel">

		<form method="POST" class="login-form" name="register" action="/user/register">
			<input id="step" name="step" type="hidden" ref="step" value="2">

			<div class="row">
				<div class="input-field col s12 center">
					<h4>Register (cont.)</h4>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<i class="mdi-social-cake prefix"></i>
					<input id="dob" name="dob" type="date" class="datepicker" ref="dob">
					<label for="dob" class="center-align">Date of Birth</label>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<input id="address" name="address" type="text" ref="address" required>
					<label for="address" class="center-align">Address</label>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<input id="city" name="city" type="text" ref="city" required>
					<label for="city" class="center-align">City</label>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s6">
					<input id="state" name="state" type="text" ref="state" required>
					<label for="state" class="center-align">State</label>
				</div>
				<div class="input-field col s6">
					<input id="country" name="country" type="text" ref="country" required>
					<label for="country" class="center-align">Country</label>
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
