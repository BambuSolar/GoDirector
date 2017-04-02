var Deploy = (function () {

    var environments = {};

    var versions = null;

    var steps = null;

    var taskId = '';

    var getVersions = function (environment, force) {

        if(force || versions == null){

            showLoader();

            var url = '/api/versions';

            $.ajax({
                    "url": url,
                    "method": 'GET'
                })
                .done(function(result, textStatus, xhr){

                    versions = result.data;

                    loadVersions(environment);

                    hideLoader();

                })
                .fail(function( jqXHR, textStatus ){

                    hideLoader();

                    showErrorMessage('An error occurred', 'Please, try it again');

                });

        }else{

            loadVersions(environment);

        }

    };

    var loadVersions = function (environment) {

        $('#deployVersionSelect').empty();

        $.each(versions[environment], function (index, item) {

            $('#deployVersionSelect')
                .append('<option value="' + item + '">' + item + '</option>');

        });

    };

    var getEnvironments = function () {

        showLoader();

        var url = '/api/environments';

        $.ajax({
                "url": url,
                "method": 'GET'
            })
            .done(function(result, textStatus, xhr){

                $('#deployEnvironmentSelect')
                    .empty()
                    .append('<option value=""> -- Select an environment -- </option>');

                $.each(result.data, function (index, item) {

                    environments[item.Name] = item;

                    $('#deployEnvironmentSelect').append(
                        '<option value="' + item.Name + '">' + item.LongName + '</option>'
                    );

                });

                hideLoader();

            })
            .fail(function( jqXHR, textStatus ){

                hideLoader();

                showErrorMessage('An error occurred', 'Please, try it again');

            });

    };

    var createSteps = function (task, environment, steps){

        $('#progress-panel').find('.deploy-steps-list').empty();

        $.each(steps[environment], function (index, item) {

            var class_step = '';

            if(task.CurrentStep == index + 1){

                if(task.Status == 'in_progress'){
                    class_step = 'list-group-item-info';
                }else{
                    if(task.Status == "error"){
                        class_step = 'list-group-item-danger';

                        if(taskId != ''){
                            swal("Error", "Deploy creation error", "error");
                            taskId = '';
                        }

                    }else{
                        class_step = 'list-group-item-success';

                        if(taskId != ''){
                            swal("Success", "Deploy successfully created", "success");
                            taskId = '';
                        }

                    }
                }

            }else{

                if(index + 1 < task.CurrentStep) {

                    class_step = 'list-group-item-success';

                }

            }

            var step = '<li class="list-group-item deploy-step ' + class_step + '">';
            step += '<div class="deploy-step-number">' + ( index + 1 )+ '</div>';
            step += '<h4 class="list-group-item-heading">' + item + '</h4>';

            if((task.CurrentStep == index + 1) && (task.Status == 'in_progress')){
                step += '<div class="progress"><div class="progress-bar progress-bar-striped active" role="progressbar" aria-valuenow="45" aria-valuemin="0" aria-valuemax="100" style="width: 100%"></div></div>';
            }

            step += '</li>';

            $('#progress-panel').find('.deploy-steps-list').append(step);

        });

        showSteps();

    };

    var listenerFormSubmit = function () {

        $('#form-deploy').validate({

            rules: {

                Version: {
                    required: true
                },
                Environment: {
                    required: true
                }

            },

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

                var data = JSON.stringify(_getFormData(form));

                var method = "POST";

                var url = "/api/deploys";

                showLoader();

                $.ajax({
                        "url": url,
                        "method": method,
                        "data": data,
                        "contentType": "application/json"
                    })
                    .done(function(result, textStatus, xhr){

                        if(xhr.status == 201 && result.success && result.data.new_task){

                            swal("Info", "Deploy created", "info");

                            taskId =  result.data.task.Id;

                        }

                        hideLoader();

                        checkRunningTask();

                    })
                    .fail(function( jqXHR ){

                        hideLoader();

                        if(jqXHR.status == 409){

                            swal("Conflict", jqXHR.responseJSON.error, "warning");

                        }else{
                            showErrorMessage('An error occurred', 'Please, try it again');
                        }

                    });
            }

        });

    };

    var _getFormData = function (form){
        var unindexed_array = $(form).serializeArray();
        var indexed_array = {};

        $.map(unindexed_array, function(n, i){
            indexed_array[n['name']] = n['value'];
        });

        return indexed_array;
    };

    var listeners = function () {

        $('#deployEnvironmentSelect').on('change', function(e){

            var env = $(this).val();
            
            if( env != 0){
                getVersions(env, false);
            }

        });

    };

    var checkRunningTask = function () {

        showForm();

        var url = "/api/deploys/last";

        $.ajax({
                "url": url,
                "method": "GET",
                "contentType": "application/json"
            })
            .done(function(result, textStatus, xhr){

                if(Object.keys(result.data.task).length > 0){

                    var task = result.data.task;

                    var environment = result.data.deploy.Environment;

                    if (task.Status != "in_progress"){
                        showForm();
                    }else{
                        hideForm();
                    }

                    if(steps == null) {

                        var url = "/api/deploys/steps";

                        $.ajax({
                                "url": url,
                                "method": "GET",
                                "contentType": "application/json"
                            })
                            .done(function(result){

                                steps = result.data;

                                createSteps(task, environment, steps);

                                hideLoader();

                            })
                            .fail(function( ){

                                hideLoader();

                                showErrorMessage('An error occurred', 'Please, try it again');

                            });

                    }else{
                        createSteps(task, environment, steps);

                        hideLoader();
                    }

                }else{
                    hideSteps();
                }

            })
            .fail(function( jqXHR, textStatus ){

                hideLoader();

                showErrorMessage('An error occurred', 'Please, try it again');

            });

    };

    var init = function () {

        checkRunningTask();

        setInterval(function(){ checkRunningTask(); }, 20 * 1000);

        getEnvironments();

        listenerFormSubmit();

        listeners();

    };

    var showSteps = function () {

        $('#progress-panel').show()

    };

    var hideSteps = function () {

        $('#progress-panel').hide();

    };

    var showForm = function () {

        $('#form-panel').show();

    };

    var hideForm = function () {

        $('#form-panel').hide();

    };

    return {
        init: init
    };


}());


$(function () {
    Deploy.init();
});

