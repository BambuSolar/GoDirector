<div class="layout">
    <div class="container">
        <h1 class="page-header">My Personal Information</h1>

        <div class="panel panel-default">
            <div class="panel-body">
                <form class="form-horizontal" method="post">
                    <div class="form-group">
                        <label for="inputFullName" class="">Full Name</label>
                        <input type="text" id="inputFullName" name="FullName" value="{{.FullName}}" class="form-control" placeholder="Full Name" required autofocus>
                    </div>
                    <div class="form-group">
                        <label for="inputEmail" class="">Email address</label>
                        <input type="email" id="inputEmail" name="Email" class="form-control" value="{{.Email}}" placeholder="Email address" required>
                    </div>
                    <div class="form-group">
                        <div class="row">
                            <div class="col-md-4 col-sm-4 col-xs-6">
                                <button type="submit" class="btn btn-primary btn-block">Save</button>
                            </div>
                            <div class="col-md-4 col-md-offset-4 col-sm-4 col-sm-offset-4 col-xs-6">
                                <button type="button" class="btn btn-default btn-block" data-dismiss="modal">Cancel</button>
                            </div>
                        </div>
                    </div>
                </form>

            </div>
        </div>

    </div>
</div>

{{template "scripts.tpl"}}
<script src="/static/js/users/edit.js"></script>