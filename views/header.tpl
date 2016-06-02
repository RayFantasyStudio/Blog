[[define "header"]]
<!DOCTYPE html>
<html>
<head>
    <title>[[.Title]]</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, minimum-scale=1, maximum-scale=1">
    <link rel="image_src" type="image/jpeg" href="/static/img/logo.png">
    <link href="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.css" class="ui" rel="stylesheet">
    <link href="/static/css/base.css" rel="stylesheet">
    <script src="//cdn.bootcss.com/jquery/2.2.4/jquery.js"></script>
    <script src="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.js"></script>

</head>
<body>
<div>
    <div class="ui main text container">
        <h1 class="ui header">[[.Owner]]的Blog</h1>
        <p>This example shows how to use lazy loaded images, a sticky menu, and a simple text container</p>
    </div>
    <div class="ui borderless main menu">
        <div class="ui text container">
            <div href="#" class="header item">
                [[.User]]
            </div>
            <a href="#" class="item">Blog</a>
            <a href="#" class="item">Articles</a>
            <a href="#" class="ui right floated dropdown item" tabindex="0">
                Dropdown <i class="dropdown icon"></i>
                <div class="menu transition hidden" tabindex="-1">
                    <div class="item">Link Item</div>
                    <div class="item">Link Item</div>
                    <div class="divider"></div>
                    <div class="header">Header Item</div>
                    <div class="item">
                        <i class="dropdown icon"></i>
                        Sub Menu
                        <div class="menu transition hidden">
                            <div class="item">Link Item</div>
                            <div class="item">Link Item</div>
                        </div>
                    </div>
                    <div class="item">Link Item</div>
                </div>
            </a>
        </div>
    </div>
</div>
[[end]]