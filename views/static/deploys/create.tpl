<div class="layout">
    <div class="container">
        <h1 class="page-header">Deploy Controller</h1>

        <div class="panel panel-default" id="form-panel">
            <div class="panel-body">

                <form id="form-deploy">

                    <div class="row">

                        <div class="col-sm-6 col-xs-12">
                            <div class="form-group">
                                <label class="sr-only" for="deployEnvironmentSelect">Environment</label>
                                <select id="deployEnvironmentSelect" class="form-control" name="Environment">
                                    <option value=""> -- Select an environment -- </option>
                                </select>
                            </div>
                        </div>

                        <div class="col-sm-6 col-xs-12">

                            <div class="form-group">
                                <label class="sr-only" for="deployVersionSelect">Version</label>
                                    <select id="deployVersionSelect" class="form-control select2"  name="Version">
                                        <option value=""> -- First select an environment -- </option>
                                    </select>
                            </div>

                        </div>

                        <div class="col-md-4 col-md-offset-8 col-xs-12">
                            <button type="submit" class="btn btn-primary btn-block">Deploy</button>
                        </div>

                    </div>
                </form>

            </div>
        </div>

        <div class="panel panel-default" id="progress-panel">
            <div class="panel-body">

                <ul class="list-group deploy-steps-list"></ul>

            </div>
        </div>

    </div>
</div>


{{template "scripts.tpl"}}
<script src="/static/js/deploys/create.js"></script>