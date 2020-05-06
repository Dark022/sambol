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
});

