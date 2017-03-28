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
                            <div class="input-group">
                                <select id="versionSelect" class="form-control">
                                    <option value="-1"> -- First select an environment -- </option>
                                </select>
                                <div class="input-group-addon" style="padding: 0 0 0 10px;">
                                    <button class="btn btn-default" type="button" data-toggle="modal" data-target="#myModal">
                                        <span class="glyphicon glyphicon-plus" aria-hidden="true"></span>
                                    </button>
                                </div>
                            </div>
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

                    <li class="list-group-item deploy-step list-group-item-danger">
                        <div class="deploy-step-number">
                            3
                        </div>

                        <h4 class="list-group-item-heading">Create Pull Request</h4>

                    </li>

                    <li class="list-group-item deploy-step">
                        <div class="deploy-step-number">
                            4
                        </div>

                        <h4 class="list-group-item-heading">Merge Pull Request</h4>

                    </li>

                    <li class="list-group-item deploy-step">
                        <div class="deploy-step-number">
                            5
                        </div>

                        <h4 class="list-group-item-heading">Create Release</h4>

                    </li>

                    <li class="list-group-item deploy-step">
                        <div class="deploy-step-number">
                            6
                        </div>

                        <h4 class="list-group-item-heading">Update Code</h4>

                    </li>

                    <li class="list-group-item deploy-step">
                        <div class="deploy-step-number">
                            7
                        </div>

                        <h4 class="list-group-item-heading">Deploy in environment Production</h4>

                    </li>

                    <li class="list-group-item deploy-step">
                        <div class="deploy-step-number">
                            8
                        </div>

                        <h4 class="list-group-item-heading">Test in environment Production</h4>

                    </li>

                </ul>

            </div>
        </div>

    </div>
</div>


{{template "scripts.tpl"}}
<script src="/static/js/deploys.js"></script>

<div class="modal fade" tabindex="-1" role="dialog" id="myModal">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                <h4 class="modal-title">Modal title</h4>
            </div>
            <div class="modal-body">
                <p>One fine body&hellip;</p>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                <button type="button" class="btn btn-primary">Save changes</button>
            </div>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->