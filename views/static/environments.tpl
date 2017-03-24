<div class="layout">
    <div class="container">
        <h1 class="page-header">Environments</h1>

        <div class="panel panel-default">
            <div class="panel-body">
                <div class="table-responsive">
                    <table class="table table-striped" id="index-table">
                        <thead>
                        <tr>
                            <th style="width: 20px;"></th>
                            <th>Name</th>
                            <th>Version</th>
                            <th>Branch</th>
                            <th style="width: 160px;">Operations</th>
                        </tr>
                        </thead>
                        <tbody></tbody>
                        <tfoot>
                        <tr>
                            <td colspan="4">
                                <nav aria-label="...">
                                    <div class="row pager">
                                        <div class="col-md-4 col-sm-4 col-xs-6 previous">
                                            <button class="btn btn-default btn-block"><span aria-hidden="true">&larr;</span> Previous</button>
                                        </div>
                                        <div class="col-md-4 col-md-offset-4 col-sm-4 col-sm-offset-4 col-xs-6 next">
                                            <button class="btn btn-default btn-block">Next <span aria-hidden="true">&rarr;</span></button>
                                        </div>
                                    </div>
                                </nav>
                            </td>
                        </tr>
                        <tr>
                            <td colspan="4"><button id="btnCreateItem" class="btn btn-block btn-success">Create New Environment</button></td>
                        </tr>
                        </tfoot>
                    </table>
                </div>
            </div>
        </div>

    </div>
</div>

{{template "scripts.tpl"}}

<script src="/static/js/CRUD_library.js"></script>
<script src="/static/js/environments.js"></script>
