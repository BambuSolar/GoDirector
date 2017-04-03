
function CRUD (config) {

    var self = this;

    this.config = config;

    this.elementSelected = -1;

    this.pagination = {
        limit: 10,
        offset: 0,
        total: 0
    };

    this.default_config = {
        path: '',
        listTable: 'table',
        nameEntity: '',
        rowTemplateToIndexTable: {
            numberRow: false,
            columns: []
        },
        parseFields:{}
    };

    this.init = function () {

        var rowTemplateToIndexTable = self.config.rowTemplateToIndexTable;

        delete self.config.rowTemplateToIndexTable;

        jQuery.extend(self.default_config, self.config);

        jQuery.extend(self.default_config.rowTemplateToIndexTable, rowTemplateToIndexTable);

        var table = self.default_config.listTable;

        var colspan = 2 + rowTemplateToIndexTable.columns.length;

        $(table).find('tfoot').find('td').attr('colspan', colspan);

        self.listeners();

        self.load_index();

    };

    this.listeners = function(){

        $('.pager').find('.next').on('click', function (e) {

            e.preventDefault();

            self._paginationNext(this);

            return false;

        });

        $('.pager').find('.previous').on('click', function (e) {

            e.preventDefault();

            self._paginationPrevious(this);

            return false;

        });

        $('#btnCreateItem').on('click', function (e) {

            e.preventDefault();

            self._createItem(this);

            return false;

        });

    };

    this._listenersTableOperations = function () {

        $('.show-item').on('click', function (e) {

            e.preventDefault();

            self._showItem(this);

            return false;

        });

        $('.edit-item').on('click', function (e) {

            e.preventDefault();

            self._editItem(this);

            return false;

        });

        $('.delete-item').on('click', function (e) {

            e.preventDefault();

            self._deleteItem(this);

            return false;

        });

    };

    this.load_index = function () {

        showLoader();

        var url = self.default_config.path;

        var pagination = 'limit=' + self.pagination.limit + '&offset=' + self.pagination.offset;

        url += '?' + pagination;

        $.ajax(url)
            .done(function(result){

                self._addRowToIndexTable(result.data);

                self._confPagination(result.pagination);

                hideLoader();

            })
            .fail(function(data){

                showErrorMessage('An error occurred', 'Please, try it again');

                hideLoader();

            });
    };

    this._createRowTemplateToIndexTable = function (index, item) {

        var row = '<tr>';

        var config = self.default_config.rowTemplateToIndexTable;

        row += '<td>' + (index + 1) + '</td>';

        var parseFields = self.default_config.parseFields;

        $.each(config.columns, function ( _ , c) {

            var value = item[c];

            if(parseFields[c]){
                if(parseFields[c]["value"]){
                    value = parseFields[c]["value"](value);
                }
            }

            row += '<td>' +value + '</td>';

        });

        var operations = '<td>';

        operations += '<div class="btn-group" role="group">';

        if(self.default_config.operations.includes("show")){

            operations += '<button class="btn btn-info show-item" type="button" data-id="' + item.Id + '"><span class="glyphicon glyphicon-eye-open" aria-hidden="true"></span></button> ';

        }

        if(self.default_config.operations.includes("edit")){

            operations += '<button class="btn btn-primary edit-item" type="button" data-id="' + item.Id + '"><span class="glyphicon glyphicon-pencil" aria-hidden="true"></span></button> ';

        }

        if(self.default_config.operations.includes("delete")){

            operations += '<button class="btn btn-danger delete-item" type="button" data-id="' + item.Id + '"><span class="glyphicon glyphicon-trash" aria-hidden="true"></span></button> ';

        }

        operations += '</div></td>';

        row += operations;

        row += '</tr>';

        return row;

    };

    this._addRowToIndexTable = function (data) {

        var table = self.default_config.listTable;

        $(table).find('tbody').empty();

        $.each(data, function (index, item) {

            var  row = self._createRowTemplateToIndexTable(index, item);

            $(table).find('tbody').append(row);

        });

        self._listenersTableOperations();

    };

    this._confPagination = function(dataPagination){

        self.pagination = dataPagination;

        var previousBtn = $('.pager').find('.previous');

        var nextBtn = $('.pager').find('.next');

        if(self.pagination.offset == 0){

            $(previousBtn).addClass('disabled');

            $(previousBtn).find('button').addClass('disabled');

        }else{

            $(previousBtn).removeClass('disabled');

            $(previousBtn).find('button').removeClass('disabled');

        }

        if( (self.pagination.offset + self.pagination.limit) >= self.pagination.total){

            $(nextBtn).addClass('disabled');

            $(nextBtn).find('button').addClass('disabled');

        }else{

            $(nextBtn).removeClass('disabled');

            $(nextBtn).find('button').removeClass('disabled');

        }

    };

    this._paginationNext = function (btn) {

        if(! $(btn).hasClass('disabled')){

            self.pagination.offset += self.pagination.limit;

            self.load_index();

        }

    };

    this._paginationPrevious = function (btn) {

        if(! $(btn).hasClass('disabled')){

            self.pagination.offset -= self.pagination.limit;

            if(self.pagination.offset < 0){

                self.pagination.offset = 0;

            }

            self.load_index();

        }

    };

    this._deleteItem = function (element) {

        self.elementSelected = $(element).attr('data-id');

        swal({
                title: "Confirm delete",
                text: "Are you sure you want to delete this item?",
                type: "warning",
                showCancelButton: true,
                confirmButtonClass: "btn-danger",
                confirmButtonText: "Delete",
                closeOnConfirm: false
            },
            function(){

                var url = self.default_config.path + '/' + self.elementSelected;

                showLoader();

                $.ajax(url,{
                        method: "DELETE"
                    })
                    .done(function(result){

                        hideLoader();

                        swal({
                            title: "Deleted",
                            text: "Item has been deleted",
                            type: "success",
                            showCancelButton: false,
                            closeOnConfirm: true
                        }, function () {
                            self.load_index();
                        });

                    })
                    .fail(function(data){

                        showErrorMessage('An error occurred', "Item, can't be deleted. Please, try it again");

                        window.sweetAlert.close();

                        hideLoader();

                    });
            });

    };

    this._getItem = function (successCallback, failCallback) {

        var url = self.default_config.path + '/' + self.elementSelected;

        showLoader();

        $.ajax(url,{
                method: "GET"
            })
            .done(function(result){

                hideLoader();

                if(successCallback){
                    successCallback(result.data);
                }

            })
            .fail(function(data){

                showErrorMessage('An error occurred', "Item, can't be deleted. Please, try it again");

                hideLoader();

                if(failCallback){
                    failCallback();
                }

            });

    };

    this._createModalShow = function (options, callback) {

        var modal = '';

        var modalElement = $('#modal-show');

        if($(modalElement).length > 0){

            $(modalElement).find('.modal-title').text(options.title);

            $(modalElement).find('.modal-body').html(options.body);

            if (options.footer){

                $(modalElement).find('.modal-footer').html(options.footer);
            }

        }else {

            modal += '<div class="modal fade" tabindex="-1" role="dialog" id="modal-show">';
            modal += '<div class="modal-dialog" role="document">';
            modal += '<div class="modal-content">';
            modal += '<div class="modal-header">';
            modal += '<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>';
            modal += '<h4 class="modal-title">' + options.title + '</h4>';
            modal += '</div>';
            modal += '<div class="modal-body">';
            modal += options.body;
            modal += '</div>';
            if (options.footer){
                modal += '<div class="modal-footer">';
                modal += options.footer;
                modal += '</div>';
            }
            modal += '</div><!-- /.modal-content -->';
            modal += '</div><!-- /.modal-dialog -->';
            modal += '</div><!-- /.modal -->';

            $('body').append(modal);

        }

        if(callback){
            callback(modal);
        }

    };

    this._showItem = function (element) {

        self.elementSelected = $(element).attr('data-id');

        var parseFields = self.default_config.parseFields;

        self._getItem(function(data){

            var body = '';

            body += '<ul class="list-group">';
            for (var key in data) {
                if (key != 'Id') {

                    var label = key;

                    var value = data[key];

                    if(parseFields[key]){
                        if(parseFields[key]["key"]){
                            label = parseFields[key]["key"](key);
                        }
                        if(parseFields[key]["value"]){
                            value = parseFields[key]["value"](value);
                        }
                    }
                    body += '<li class="list-group-item"><strong>' + label + '</strong>: ' + value + '</li>';
                }
            }
            body += '</ul>';

            var options = {
                title: 'Show ' + self.default_config.nameEntity,
                body: body
            };

            self._createModalShow(options, function (modal) {

                $('#modal-show').modal('show');

            });

        });

    };

    this._editItem = function (element) {

        self.elementSelected = $(element).attr('data-id');

        self._getItem(function(data){

            var options = {
                title: 'Edit ' + self.default_config.nameEntity,
                data: data
            };

            self._createModalForm(options, function () {

                $('#modal-form').find('form').attr('method', 'PUT');

                var url = self.default_config.path + '/' + self.elementSelected;

                $('#modal-form').find('form').attr('action', url);


                $('#modal-form').modal('show');

            });

        });

    };

    this._createItem = function (element) {

        var options = {
            title: 'Create ' + self.default_config.nameEntity
        };

        self._createModalForm(options, function () {

            $('#modal-form').find('form').attr('method', 'POST');

            var url = self.default_config.path;

            $('#modal-form').find('form').attr('action', url);

            $('#modal-form').modal('show');

        });

    };

    this._createModalForm = function (options, callback) {

        var modal = '';

        var modalElement = $('#modal-form');

        if($(modalElement).length > 0){

            if(options.data) {

                $.each(self.default_config.formFields, function (index, item) {

                    self._setDefaultValue(item, options.data[item.field]);

                });

            }else{

                $.each(self.default_config.formFields, function(index, item){

                    body += self._setDefaultValue(item);

                });

            }

            $(modalElement).find('.form-group')
                .removeClass('has-error')
                .removeClass('has-success')
                .find('.help-block').remove();

            $(modalElement).find('.modal-title').text(options.title);

        }else {

            var body = '';

            if(options.data) {

                $.each(self.default_config.formFields, function(index, item){

                    body += self._createField(item, options.data[item.field]);

                });

            }else{

                $.each(self.default_config.formFields, function(index, item){

                    body += self._createField(item);

                });

            }

            var footer = '';

            footer += '<div class="col-md-4 col-sm-4 col-xs-6">';
            footer += '<button type="submit" class="btn btn-primary btn-block">Save</button>';
            footer += '</div>';

            footer += '<div class="col-md-4 col-md-offset-4 col-sm-4 col-sm-offset-4 col-xs-6">';
            footer += '<button type="button" class="btn btn-default btn-block" data-dismiss="modal">Cancel</button>';
            footer += '</div>';

            modal += '<div class="modal fade" tabindex="-1" role="dialog" id="modal-form">';
            modal += '<div class="modal-dialog" role="document">';
            modal += '<div class="modal-content">';
            modal += '<div class="modal-header">';
            modal += '<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>';
            modal += '<h4 class="modal-title">' + options.title + '</h4>';
            modal += '</div>';
            modal += '<form>';
            modal += '<div class="modal-body">';
            modal += body;
            modal += '</div>';
            modal += '<div class="modal-footer">';
            modal += '<div class="row">';
            modal += footer;
            modal += '</div>';
            modal += '</div>';
            modal += '</form>';
            modal += '</div><!-- /.modal-content -->';
            modal += '</div><!-- /.modal-dialog -->';
            modal += '</div><!-- /.modal -->';

            $('body').append(modal);

            self._activateForm();

        }

        if(callback){
            callback(modal);
        }

    };

    this._activateForm = function(){

        var rules = {};

        $.each(self.default_config.formFields, function (index, item) {

            var rule = item.options.rule;

            rules[item.field] = rule;
        });

        $('#modal-form').find('form').validate({
            rules: rules,
            errorPlacement: function(error, element) {

                var text = $(error).text();

                var parent = $(element).parent();

                parent.removeClass('has-success').addClass('has-error');

                if(text){

                    parent.find('.help-block').remove();

                    parent.append('<span class="help-block">' + text + '</span>');

                }
            },
            success: function(element) {

                var id = $(element).attr('id').split('-')[0];

                var parent = $('#' + id).parent();

                parent.removeClass('has-error').addClass('has-success');

                parent.find('.help-block').remove();

            },
            submitHandler: function(form) {

                var data = self.normalizeBodyRequest(JSON.stringify(self._getFormData(form)));
                
                var method = $(form).attr('method');

                var url = $(form).attr('action');

                showLoader();

                $.ajax({
                        "url": url,
                        "method": method,
                        "data": data,
                        "contentType": "application/json"
                    })
                    .done(function(result, textStatus, xhr){

                        hideLoader();

                        self._successFormRequest(result, textStatus, xhr);

                    })
                    .fail(function( jqXHR, textStatus ){

                        hideLoader();

                        self._failFormRequest(jqXHR, textStatus );

                    });

            }

        });

    };

    this.normalizeBodyRequest = function(requestBody){

        return requestBody.replace("\"true\"", "true").replace("\"false\"", "false");

    };

    this._getFormData = function (form){
        var unindexed_array = $(form).serializeArray();
        var indexed_array = {};

        $.map(unindexed_array, function(n, i){
            indexed_array[n['name']] = n['value'];
        });

        return indexed_array;
    };

    this._successFormRequest = function(result, textStatus, xhr){

        if(result.success){

            var operation = "edited";

            var operationTitle = "Edition";


            if(xhr.status == 201){

                operation = "created";

                operationTitle = "Creation";

            }

            swal({
                title: operationTitle,
                text: "Item has been " + operation + " successfully",
                type: "success",
                showCancelButton: false,
                closeOnConfirm: true
            }, function () {

                self.load_index();

                window.sweetAlert.close();

                $('#modal-form').modal('hide');

            });

        }else{

            swal({
                title: "An error occurred",
                text: result.error,
                type: "error",
                showCancelButton: false,
                closeOnConfirm: true
            }, function () {

                window.sweetAlert.close();

            });

        }


    };

    this._failFormRequest = function(jqXHR, textStatus ){
        if(jqXHR.statusText == "error"){

            showErrorMessage('An error occurred', 'Please, try it again');

        }
    };


    this._createField = function(item, value){

        var config_field = item.options.form;

        var html_field;

        switch(config_field.type.split(':')[0]){
            case 'input':

                html_field = '<div class="form-group">';
                html_field += '<label for="field'+ item.field +'">' + config_field.label + '</label>';
                html_field += '<input type="' + config_field.type.split(':')[1] + '" name="' + item.field + '" ';

                if(value){
                    html_field += 'value="' + value + '" ';
                }

                if(config_field.autocomplete =! undefined && !config_field.autocomplete){
                    html_field += 'autocomplete="off" ';
                }

                html_field += 'class="form-control" id="field' + item.field + '" placeholder="' + config_field.placeholder + '">';
                html_field += '</div>';

                break;
            case 'checkbox':


                html_field = '<div class="form-group">';

                html_field += '<label for="field'+ item.field +'">' + config_field.label + '</label>';
                html_field += '</div>';
                html_field += '<div class="radio">';
                html_field += '<label>';
                html_field += '<input type="radio" name="' + item.field + '" value="true" checked>';
                html_field += 'Yes';
                html_field += '</label>';
                html_field += '</div>';
                html_field += '<div class="radio">';
                html_field += '<label>';
                html_field += '<input type="radio" name="' + item.field + '" value="false">';
                html_field += 'No';
                html_field += '</label>';
                html_field += '</div>';

                break;
        }

        return html_field;

    };

    this._setDefaultValue = function (item, data) {

        switch(item.options.form.type.split(':')[0]) {
            case 'input':

                if (!data) data = '';

                $("#field" + item.field).val(data);

                break;
        }

    };

}
