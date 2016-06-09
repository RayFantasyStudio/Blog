[[define "paginator"]]
<div class="ui pagination menu">
    [[range .]]
    <a class="[[if .IsActive]]active[[end]] [[if .IsDisabled]]disabled[[end]] item" [[if .IsDisabled]]disabled[[end]]
       [[if .Link]]href="[[.Link]]" [[end]]>
        [[.Lable]]
    </a>
    [[end]]
</div>
[[end]]