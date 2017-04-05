<div class="layout">
    <div class="container">
        <div class="jumbotron">
            <p class="text-center">
                <img src="static/img/icon.png" style="max-width:300px"  >
            </p>
            <h1 class="text-center">GoDirector</h1>
            <p class="text-justify">
            This system allows to administer the builds in the different environments, as well as the deploys.
            The system has a friendly graphical interface that allows to easily visualize the status of the different tasks.
            In turn, the system has an internal control that ensures that no damage or unwanted effects are observed when sending orders to PythonTransformers.
            </p>
            <p class="text-center">
                <div class="row">
                    <div class="col-xs-12 col-sm-4 col-sm-offset-2">
                        <a class="btn btn-primary btn-lg btn-block m-10" href="/builds/new" role="button">Create a Build</a>
                    </div>
                    <div class="col-xs-12 col-sm-4">
                        <a class="btn btn-primary btn-lg btn-block m-10" href="/deploys/new" role="button">Create a Deploy</a>
                    </div>
                </div>
            </p>
        </div>
    </div>
</div>
{{template "scripts.tpl"}}
<script src="/static/js/index.js"></script>