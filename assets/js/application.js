require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");

$(() => {
    const saveCampaignBtn = document.getElementById("saveCampaign")

    $(saveCampaignBtn).click(function(){
        const checkboxs = $('input[type="checkbox"]');
        checkboxs.each(function(){
            if(!$(this).is(":checked")) {
                $(this).parent().parent().next().remove();     
            }
        });
    });

    $("#search").keyup(function() {
        let filter = this.value.toUpperCase();
        const templatetable = document.getElementById("templateTable");
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
});

