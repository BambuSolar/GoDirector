<div class="layout">
    <div class="container">
        <h1 class="page-header">Deploy Controller</h1>

        <div class="panel panel-default">
            <div class="panel-body">

                <div class="row">

                    <div class="col-sm-4 col-xs-12">
                        <div class="form-group">
                            <label class="sr-only" for="environmentSelect">Environment</label>
                            <select id="environmentSelect" class="form-control">
                                <option value="-1"> -- Select an environment -- </option>
                            </select>
                        </div>
                    </div>

                    <div class="col-sm-4 col-xs-12">

                        <div class="form-group">
                            <label class="sr-only" for="versionSelect">Version</label>
                            <select id="versionSelect" class="form-control">
                                <option value="-1"> -- First select an environment -- </option>
                            </select>
                        </div>

                    </div>

                    <div class="col-sm-4 col-xs-12">
                        <button type="button" class="btn btn-primary btn-block">Deploy</button>
                    </div>

                </div>

            </div>
        </div>

        <div class="panel panel-default">
            <div class="panel-body">

                <div class="progress">
                    <div class="progress-bar" role="progressbar" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100" style="width: 60%;">
                        60%
                    </div>
                </div>

                <ul class="list-group deploy-steps-list">

                    <li class="list-group-item deploy-step list-group-item-success">
                        <div class="deploy-step-number">
                            1
                        </div>

                        <h4 class="list-group-item-heading">Deploy in environment Beta</h4>

                    </li>

                    <li class="list-group-item deploy-step list-group-item-info">
                        <div class="deploy-step-number">
                            2
                        </div>

                        <h4 class="list-group-item-heading">Test in environment Beta</h4>

                    </li>

                </ul>

            </div>
        </div>

    </div>
</div>


{{template "scripts.tpl"}}
<script src="/static/js/deploys/create.js"></script>