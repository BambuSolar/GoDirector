$(function(){

    var environments = new CRUD({
        path: '/api/environments',
        listTable: '#index-table',
        nameEntity: 'System Parameter',
        rowTemplateToIndexTable: {
            columns: ['Name', 'Version', 'Branch']
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
                        label: "RootPath",
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

    environments.init()
});