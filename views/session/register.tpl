<form class="form-register" method="post">
    <p class="text-center">
        <h1 class="text-center">
            <img src="static/img/icon.png" style="max-height:100px">
        </h1>
    </p>
    <h2 class="form-register-heading text-center">Registration</h2>

    <label for="inputFullName" class="sr-only">Full Name</label>
    <input type="text" id="inputFullName" name="FullName" class="form-control" placeholder="Full Name" required autofocus>

    <label for="inputEmail" class="sr-only">Email address</label>
    <input type="email" id="inputEmail" name="Email" class="form-control" placeholder="Email address" required>

    <label for="inputPassword" class="sr-only">Password</label>
    <input type="password" id="inputPassword" name="Password" class="form-control" placeholder="Password" required>

    <label for="inputRePassword" class="sr-only">Re Password</label>
    <input type="password" id="inputRePassword" name="Repassword" class="form-control" placeholder="Re Password" required>

    <div class="g-recaptcha" data-sitekey="6Ld_exsUAAAAACfhvBW0K6Ck8DlO7YBO_Z4vbz6g"></div>

    <button class="btn btn-lg btn-primary btn-block" type="submit">Submit</button>

    Got an account? <a href="/login">Log in now!</a>

</form>

{{template "scripts.tpl"}}
<script src="/static/js/session/signup.js"></script>