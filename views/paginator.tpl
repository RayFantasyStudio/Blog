[[define "paginator"]]
<div class="ui pagination menu">
    [[range .]]
    <a class="[[if .IsActive]]active[[end]] [[if .IsDisabled]]disabled[[end]] item" href="[[.Link]]">
        [[.Lable]]
    </a>
    [[end]]
</div>
[[end]]