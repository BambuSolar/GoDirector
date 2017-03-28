<div class="layout">
    <div class="container">
        <h1 class="page-header">Build Controller</h1>

        <div class="panel panel-default" id="form-panel">
            <div class="panel-body">

                <form id="form-build">
                    <div class="form-group">
                        <label class="" for="buildEnvironmentSelect">Environment</label>
                        <select id="buildEnvironmentSelect" class="form-control" name="Environment">
                            <option value="-1"> -- Select an environment -- </option>
                        </select>
                    </div>
                    <div class="form-group">
                        <label for="buildUrlInput">Url</label>
                        <input type="url" class="form-control" id="buildUrlInput" placeholder="Url" name="Url">
                    </div>

                    <div class="row">
                        <div class="col-md-4 col-sm-4 col-xs-6">
                            <button type="submit" class="btn btn-primary btn-block">Build</button>
                        </div>
                        <div class="col-md-4 col-md-offset-4 col-sm-4 col-sm-offset-4 col-xs-6">
                            <button type="button" class="btn btn-default btn-block" data-dismiss="modal">Cancel</button>
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
<script src="/static/js/build.js"></script>