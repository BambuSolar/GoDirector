var Build = (function () {

    var environments = {};

    var steps = [];

    var taskId = '';

    var getEnvironments = function () {

        showLoader();

        var url = '/api/environments';

        $.ajax({
            "url": url,
            "method": 'GET'
        })
            .done(function(result, textStatus, xhr){

                $('#buildEnvironmentSelect')
                    .empty();

                $.each(result.data, function (index, item) {

                    environments[item.Name] = item;

                    $('#buildEnvironmentSelect').append(
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

    var listenerFormSubmit = function () {

        $('#form-build').validate({

            rules: {

                Url: {
                    required: true,
                    url: true
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

                var url = "/api/builds";

                showLoader();

                $.ajax({
                    "url": url,
                    "method": method,
                    "data": data,
                    "contentType": "application/json"
                })
                    .done(function(result, textStatus, xhr){

                        if(xhr.status == 201 && result.success && result.data.new_task){

                            swal("Info", "Build created", "info");

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


    };

    var checkRunningTask = function () {

        showForm();

        var url = "/api/builds/last";

        $.ajax({
            "url": url,
            "method": "GET",
            "contentType": "application/json"
        })
            .done(function(result, textStatus, xhr){

                if(result.data.length > 0){

                    var task = result.data[0];

                    if (task.Status != "in_progress"){
                        showForm();
                    }else{
                        hideForm();
                    }

                    if(steps.length < 1){

                        var url = "/api/builds/steps";

                        $.ajax({
                            "url": url,
                            "method": "GET",
                            "contentType": "application/json"
                        })
                            .done(function(result){

                                steps = result.data;

                                createSteps(task, steps);

                                hideLoader();

                            })
                            .fail(function( ){

                                hideLoader();

                                showErrorMessage('An error occurred', 'Please, try it again');

                            });

                    }else{
                        createSteps(task, steps);

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

    var createSteps = function (task, steps){

        $('#progress-panel').find('.build-steps-list').empty();

        $.each(steps, function (index, item) {

            var class_step = '';

            if(task.CurrentStep == index + 1){

                if(task.Status == 'in_progress'){
                    class_step = 'list-group-item-info';
                }else{
                    if(task.Status == "error"){
                        class_step = 'list-group-item-danger';

                        if(taskId != ''){
                            swal("Error", "Build creation error", "error");
                            taskId = '';
                        }

                    }else{
                        class_step = 'list-group-item-success';

                        if(taskId != ''){
                            swal("Success", "Build successfully created", "success");
                            taskId = '';
                        }

                    }
                }

            }else{

                if(index + 1 < task.CurrentStep) {

                    class_step = 'list-group-item-success';

                }

            }

            var step = '<li class="list-group-item build-step ' + class_step + '">';
            step += '<div class="build-step-number">' + ( index + 1 )+ '</div>';
            step += '<h4 class="list-group-item-heading">' + item + '</h4>';

            if((task.CurrentStep == index + 1) && (task.Status == 'in_progress')){
                step += '<div class="progress"><div class="progress-bar progress-bar-striped active" role="progressbar" aria-valuenow="45" aria-valuemin="0" aria-valuemax="100" style="width: 100%"></div></div>';
            }

            step += '</li>';

            $('#progress-panel').find('.build-steps-list').append(step);

        });

        showSteps();

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

    var init = function () {

        checkRunningTask();

        setInterval(function(){ checkRunningTask(); }, 20 * 1000);

        getEnvironments();

        listenerFormSubmit();

        listeners();

    };

    return {
        init: init
    };


}());


$(function () {
    Build.init();
});

