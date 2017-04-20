<div class="layout">
    <div class="container">
        <h1 class="page-header">Build Controller</h1>

        <div class="panel panel-default" id="form-panel">
            <div class="panel-body">

                <form id="form-build">
                    <div class="row">
                        <div class="col-sm-6 col-xs-12">
                            <div class="form-group">
                                <select id="buildEnvironmentSelect" class="form-control" name="Environment">
                                    <option value="-1"> -- Select an environment -- </option>
                                </select>
                            </div>
                        </div>
                        <div class="col-sm-6 col-xs-12">
                            <button type="submit" class="btn btn-primary btn-block">Build</button>
                        </div>
                    </div>
                </form>

            </div>
        </div>

        <div class="panel panel-default" id="progress-panel">
            <div class="panel-body">

                <ul class="list-group build-steps-list"></ul>

            </div>
        </div>

    </div>
</div>


{{template "scripts.tpl"}}
<script src="/static/js/builds/create.js"></script>