$(function(){

    var environments = new CRUD({
        path: '/api/environments',
        listTable: '#index-table',
        nameEntity: 'System Parameter',
        operations: ['show','edit', 'delete', 'create'],
        rowTemplateToIndexTable: {
            columns: ['Name', 'Version', 'Branch']
        },
        parseFields:{
            "Name": {
                "value": function (env) {
                    return toTitleCase(env);
                }
            },
            "Branch": {
                "value": function (env) {
                    return toTitleCase(env);
                }
            },
            "BuddyPipelineId":{
                "key": function () {
                    return "Buddy Pipeline";
                }
            },
            "AllowDirectDeploy":{
                "key": function () {
                    return "Allow Direct Deploy";
                }
            },
            "ServerUrl":{
                "key": function () {
                    return "Server";
                },
                "value": function (url) {

                    if(!url.includes("http://")){
                        url = "http://" + url;
                    }

                    return '<a href="' + url + '" target="_blank">' + url + '</a>';
                }
            },
            "UserFTP":{
                "key": function () {
                    return "User";
                }
            },
            "PasswordFTP":{
                "key": function () {
                    return "Password";
                }
            }
        },
        formFields: [
            {
                field: "Name",
                options: {
                    form: {
                        type: "input:text",
                        label: "Name",
                        placeholder: "Name",
                        autocomplete: false
                    },
                    rule: {
                        required: true,
                        minlength: 4
                    }
                }
            },
            {
                field: "LongName",
                options: {
                    form: {
                        type: "input:text",
                        label: "Long Name",
                        placeholder: "Long Name",
                        autocomplete: false
                    },
                    rule: {
                        required: true,
                        minlength: 4
                    }
                }
            },
            {
                field: "Version",
                options: {
                    form: {
                        type: "input:text",
                        label: "Version",
                        placeholder: "0.0.1",
                        autocomplete: false
                    },
                    rule: {
                        required: true,
                        minlength: 5
                    }
                }
            },
            {
                field: "Branch",
                options: {
                    form: {
                        type: "input:text",
                        label: "Branch",
                        placeholder: "Branch",
                        autocomplete: false
                    },
                    rule: {
                        required: true,
                        minlength: 4
                    }
                }
            },
            {
                field: "BuddyPipelineId",
                options: {
                    form: {
                        type: "input:text",
                        label: "Buddy Pipeline",
                        placeholder: "Buddy Pipeline",
                        autocomplete: false
                    },
                    rule: {
                        required: true
                    }
                }
            },
            {
                field: "AllowDirectDeploy",
                options: {
                    form: {
                        type: "checkbox",
                        label: "Allow Direct Deploy"
                    }
                }
            },
            {
                field: "ServerUrl",
                options: {
                    form: {
                        type: "input:text",
                        label: "Server",
                        placeholder: "Server",
                        autocomplete: false
                    },
                    rule: {
                        required: true,
                        minlength: 4
                    }
                }
            },
            {
                field: "FTPRootPath",
                options: {
                    form: {
                        type: "input:text",
                        label: "Root Path",
                        placeholder: "RootPath",
                        autocomplete: false
                    },
                    rule: {
                        required: true,
                        minlength: 4
                    }
                }
            },
            {
                field: "UserFTP",
                options: {
                    form: {
                        type: "input:text",
                        label: "User",
                        placeholder: "User",
                        autocomplete: false
                    },
                    rule: {
                        required: true,
                        minlength: 4
                    }
                }
            },
            {
                field: "PasswordFTP",
                options: {
                    form: {
                        type: "input:password",
                        label: "Password",
                        placeholder: "Password",
                        autocomplete: false
                    },
                    rule: {
                        required: true,
                        minlength: 8
                    }
                }
            }
        ]
    });

    environments.init();
});