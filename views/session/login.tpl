<form class="form-signin" method="post">

    <p class="text-center">
        <h1 class="text-center">
            <img src="static/img/icon.png" style="max-height:100px">
        </h1>
    </p>

    <h2 class="form-signin-heading">Please sign in</h2>
    <label for="inputEmail" class="sr-only">Email address</label>
    <input type="email" id="inputEmail" name="Email" class="form-control" placeholder="Email address" required autofocus>

    <label for="inputPassword" class="sr-only">Password</label>
    <input type="password" id="inputPassword" name="Password" class="form-control" placeholder="Password" required>

    <div class="g-recaptcha" data-sitekey="6Ld_exsUAAAAACfhvBW0K6Ck8DlO7YBO_Z4vbz6g"></div>

    <button class="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>

    Need an account? <a href="/signup">Sign up.</a>

</form>

{{template "scripts.tpl"}}
<script src="/static/js/session/login.js"></script>