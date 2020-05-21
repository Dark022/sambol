require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");

$(() => {
    const saveCampaignBtn = document.getElementById("saveCampaign")
    $(saveCampaignBtn).click(function () {
        const checkboxs = $('input[type="checkbox"]');
        checkboxs.each(function () {
            if (!$(this).is(":checked")) {
                $(this).parent().parent().next().remove();
            }
        });
    });

    $("#search").keyup(function () {
        let filter = this.value.toUpperCase();
        const templatetable = document.getElementById("table");
        const tr = templatetable.getElementsByTagName("tr");

        console.log(this, filter, tr)

        for (i = 0; i < tr.length; i++) {
            td = tr[i].getElementsByTagName("td")[0];
            if (td) {
                txtValue = td.textContent || td.innerText;
                if (txtValue.toUpperCase().indexOf(filter) > -1) {
                    tr[i].style.display = "";
                } else {
                    tr[i].style.display = "none";
                }
            }
        }
    });

    $('#campaign-TemplateID option:selected').each(function () {
            $.ajax({
                url: '',
                type: 'get',
                dataType: "json",
                data: {templateID: this.value},
                success: function(data) { 
                    $('#userListing').empty()             
                    for(var i = 0; i < data.length; i++) {
                        const userList =
                        `
                        <div class="col-12">
                            <div class="form-group">
                                <label>
                                    <input class="" id="campaign-UserCheck" name="UserCheck" type="checkbox" value="true">
                                    ${data[i].firstname} ${data[i].lastname}
                                </label>
                            </div>
                            <input id="campaign-UsersID" name="UsersID" tags-field="UsersID" type="hidden" value="${data[i].id}">                 
                        </div>                         
                        `

                        $('#userListing').append(userList)
                    }
                },
            });
           
    });

    $("#campaign-TemplateID").change(function () {
        let selected = $(this).children("option:selected").val();
        if (selected.value != ""){
            $.ajax({
                url: '',
                type: 'get',
                dataType: "json",
                data: {templateID: selected},
                success: function(data) {
                    $('#userListing').empty()
                    for(var i = 0; i < data.length; i++) {
                        const userList =
                        `
                        <div class="col-12">
                            <div class="form-group">
                                <label>
                                    <input class="" id="campaign-UserCheck" name="UserCheck" type="checkbox" value="true">
                                    ${data[i].firstname} ${data[i].lastname}
                                </label>
                            </div>
                            <input id="campaign-UsersID" name="UsersID" tags-field="UsersID" type="hidden" value="${data[i].id}">                 
                        </div>                         
                        `

                        $('#userListing').append(userList)
                    }
                }
            });
        }     
    });

    $('#campaign-TemplateID option:selected').each(function () {
        $.ajax({
            url: 'save',
            type: 'get',
            dataType: "json",
            data: {templateID: this.value},
            success: function(data) { 
                $('#userListing').empty()             
                for(var i = 0; i < data.length; i++) {
                    const userList =
                    `
                    <div class="col-12">
                        <div class="form-group">
                            <label>
                                <input class="" id="campaign-UserCheck" name="UserCheck" type="checkbox" value="true">
                                ${data[i].firstname} ${data[i].lastname}
                            </label>
                        </div>
                        <input id="campaign-UsersID" name="UsersID" tags-field="UsersID" type="hidden" value="${data[i].id}">                 
                    </div>                         
                    `

                    $('#userListing').append(userList)
                }
            },
        });
       
    });

    $("#campaign-TemplateID").change(function () {
        let selected = $(this).children("option:selected").val();
        if (selected.value != ""){
            $.ajax({
                url: 'save',
                type: 'get',
                dataType: "json",
                data: {templateID: selected},
                success: function(data) {
                    $('#userListing').empty()
                    for(var i = 0; i < data.length; i++) {
                        const userList =
                        `
                        <div class="col-12">
                            <div class="form-group">
                                <label>
                                    <input class="" id="campaign-UserCheck" name="UserCheck" type="checkbox" value="true">
                                    ${data[i].firstname} ${data[i].lastname}
                                </label>
                            </div>
                            <input id="campaign-UsersID" name="UsersID" tags-field="UsersID" type="hidden" value="${data[i].id}">                 
                        </div>                         
                        `

                        $('#userListing').append(userList)
                    }
                }
            });
        }     
    });
});