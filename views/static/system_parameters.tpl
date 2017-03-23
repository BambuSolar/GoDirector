<div class="layout">
    <div class="container">
        <h1 class="page-header">Sistem Parameters</h1>

        <div class="panel panel-default">
            <div class="panel-body">
                <div class="table-responsive">
                    <table class="table table-striped table-hover" id="index-table">
                        <thead>
                        <tr>
                            <th style="width: 20px;"></th>
                            <th>Key</th>
                            <th>Value</th>
                            <th style="width: 160px;">Operations</th>
                        </tr>
                        </thead>
                        <tbody></tbody>
                        <tfoot>
                            <tr>
                                <td colspan="4">
                                    <nav aria-label="...">
                                        <ul class="pager">
                                            <li class="previous">
                                                <button class="btn btn-default"><span aria-hidden="true">&larr;</span> Previous</button>
                                            </li>
                                            <li class="next">
                                                <button class="btn btn-default">Next <span aria-hidden="true">&rarr;</span></button>
                                            </li>
                                        </ul>
                                    </nav>
                                </td>
                            </tr>
                            <tr>
                                <td colspan="4"><button id="btnCreateItem" class="btn btn-block btn-success">Create New System Parameter</button></td>
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
<script src="/static/js/system_parameters.js"></script>
