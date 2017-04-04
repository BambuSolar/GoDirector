<!DOCTYPE html>
<html>
<head>
    <title>GoDirector</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link href="/static/img/icon.png" rel="icon">
    <link rel="stylesheet" href="../static/css/bootstrap.min.css">
    <link rel="stylesheet" href="../static/css/sweetalert.css">
    <link rel="stylesheet" href="../static/css/style.css">
</head>
<body>

<nav class="navbar navbar-default navbar-fixed-top">
    <div class="container">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="/" style="min-width: 200px">
                <img alt="Brand" src="static/img/icon.png" height="20px" style="float: left; margin-right: 10px; display: inline;">
                GoDirector
            </a>
        </div>
        <div id="navbar" class="navbar-collapse collapse">
            <ul class="nav navbar-nav">
                <li class="dropdown">
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Deploys <span class="caret"></span></a>
                    <ul class="dropdown-menu">
                        <li><a href="/deploys">List Deploys</a></li>
                        <li><a href="/deploys/new">Create Deploy</a></li>
                    </ul>
                </li>
                <li class="dropdown">
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Builds <span class="caret"></span></a>
                    <ul class="dropdown-menu">
                        <li><a href="/builds">List Builds</a></li>
                        <li><a href="/builds/new">Create Build</a></li>
                    </ul>
                </li>
                <li class="dropdown">
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Configurations <span class="caret"></span></a>
                    <ul class="dropdown-menu">
                        <li><a href="/system_parameters">System Parameters</a></li>
                        <li><a href="/environments">Environments</a></li>
                        <li><a href="/users">Users</a></li>
                        <li><a href="/applications">Applications</a></li>
                    </ul>
                </li>
            </ul>
            <ul class="nav navbar-nav navbar-right">
                <li class="dropdown">
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
                        <span class="glyphicon glyphicon-user" aria-hidden="true"></span>
                        {{.Userinfo.FullName}} <span class="caret"></span>
                    </a>
                    <ul class="dropdown-menu">
                        <li>
                            <a href="#">
                                <span class="glyphicon glyphicon-pencil" aria-hidden="true"></span> Modify User Data
                            </a>
                        </li>
                        <li>
                            <a href="#">
                                <span class="glyphicon glyphicon-pencil" aria-hidden="true"></span> Change Password
                            </a>
                        </li>
                        <li role="separator" class="divider"></li>
                        <li>
                            <a href="/logout">
                                <span class="glyphicon glyphicon-log-out" aria-hidden="true"></span> Logout
                            </a>
                        </li>
                    </ul>
                </li>
            </ul>
        </div><!--/.nav-collapse -->
    </div>
</nav>

{{.LayoutContent}}

<div class="loader">
    <div class="sk-cube-grid">
        <div class="sk-cube sk-cube1"></div>
        <div class="sk-cube sk-cube2"></div>
        <div class="sk-cube sk-cube3"></div>
        <div class="sk-cube sk-cube4"></div>
        <div class="sk-cube sk-cube5"></div>
        <div class="sk-cube sk-cube6"></div>
        <div class="sk-cube sk-cube7"></div>
        <div class="sk-cube sk-cube8"></div>
        <div class="sk-cube sk-cube9"></div>
    </div>
</div>

</body>
</html>