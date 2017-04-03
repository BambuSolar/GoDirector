$(function () {
    var builds = new CRUD({
        path: '/api/builds',
        listTable: '#index-table',
        nameEntity: 'builds',
        operations: ['show'],
        rowTemplateToIndexTable: {
            columns: ['Environment', 'Status', 'CreateAt']
        },
        parseFields:{
            "Environment": {
                "value": function (env) {
                    return toTitleCase(env);
                }
            },
            "Status":{
                "value": function (status) {
                    switch(status){
                        case "successful":
                            return '<span class="glyphicon glyphicon-ok" style="color: darkgreen" aria-hidden="true"></span>';
                        case "failed":
                            return '<span class="glyphicon glyphicon-remove" style="color: darkred" aria-hidden="true"></span>';
                        case "in_progress":
                            return 'Deploy in progress<img src="static/img/animations/in_progress.gif" height="24" style="margin-left: 10px;">';
                    }
                }
            },
            "ReleaseIdGitHub": {
                "key": function () {
                    return "GitHub Release"
                },
                "value": function (id) {
                    if(id == ""){
                        return " Without release ";
                    }else{
                        return id;
                    }
                }
            },
            "CreateAt":{
                "key": function () {
                    return "Creation Date";
                },
                "value": function (CreateAt) {
                    moment(CreateAt, "YYYY-MM-DDTHH:mm:ss.SSSZ");
                    return moment(CreateAt, "YYYY-MM-DDTHH:mm:ss.SSSZ").format("DD/MM/YYYY - hh:mm:ss a");
                }
            }
        }
    });

    builds.init();

});