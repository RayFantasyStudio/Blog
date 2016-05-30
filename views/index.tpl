<!DOCTYPE html>

<html>
<head>
    <title>[[.Owner]]的Blog</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, minimum-scale=1, maximum-scale=1">
    <link rel="image_src" type="image/jpeg" href="/static/img/logo.png">
    <link href="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.css" class="ui" rel="stylesheet">
    <link href="/static/css/base.css" rel="stylesheet">
    <script src="//cdn.bootcss.com/jquery/2.2.4/jquery.js"></script>
    <script src="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.js"></script>
    <script>
        $(document)
                .ready(function () {

                    // fix main menu to page on passing
                    $('.main.menu').visibility({
                        type: 'fixed'
                    });
                    $('.overlay').visibility({
                        type: 'fixed',
                        offset: 80
                    });

                    // lazy load images
                    $('.image').visibility({
                        type: 'image',
                        transition: 'vertical flip in',
                        duration: 500
                    });

                    // show dropdown on hover
                    $('.main.menu  .ui.dropdown').dropdown({
                        on: 'hover'
                    });
                })
        ;
    </script>

</head>

<body>

<div class="ui main text container">
    <h1 class="ui header">[[.Owner]]的Blog</h1>
    <p>This example shows how to use lazy loaded images, a sticky menu, and a simple text container</p>
</div>

<div class="ui borderless main menu">
    <div class="ui text container">

            <a class="item">
                Home
            </a>
            <div class="ui pointing dropdown link item">
                <span class="text">Shopping</span>
                <i class="dropdown icon"></i>
                <div class="menu">
                    <div class="header">Categories</div>
                    <div class="item">
                        <i class="dropdown icon"></i>
                        <span class="text">Clothing</span>
                        <div class="menu">
                            <div class="header">Mens</div>
                            <div class="item">Shirts</div>
                            <div class="item">Pants</div>
                            <div class="item">Jeans</div>
                            <div class="item">Shoes</div>
                            <div class="divider"></div>
                            <div class="header">Womens</div>
                            <div class="item">Dresses</div>
                            <div class="item">Shoes</div>
                            <div class="item">Bags</div>
                        </div>
                    </div>
                    <div class="item">Home Goods</div>
                    <div class="item">Bedroom</div>
                    <div class="divider"></div>
                    <div class="header">Order</div>
                    <div class="item">Status</div>
                    <div class="item">Cancellations</div>
                </div>
            </div>
            <a class="item">
                Forums
            </a>
            <a class="item">
                Contact Us
            </a>

    </div>
</div>

<div class="ui text container">
    <div class="ui relaxed divided items">
        [[range $index, $article := .Articles]]
        <div class="item">
            <div class="content">
                <a class="header">[[$article.Title]]</a>
                <div class="meta">
                    <a>[[$article.Created | SinceTime]]</a>
                    <a>Category</a>
                </div>
                <div class="description">
                    [[$article.Content]]
                </div>
                <div class="extra">
                    <div class="ui right floated primary button">
                        详情
                        <i class="right chevron icon"></i>
                    </div>
                    <div class="ui label">Tag</div>
                </div>
            </div>
        </div>
        [[end]]
    </div>
</div>

<div class="ui inverted vertical footer segment">
    <div class="ui center aligned container">
        <div class="ui stackable inverted divided grid">
            <div class="three wide column">
                <h4 class="ui inverted header">Group 1</h4>
                <div class="ui inverted link list">
                    <a href="#" class="item">Link One</a>
                    <a href="#" class="item">Link Two</a>
                    <a href="#" class="item">Link Three</a>
                    <a href="#" class="item">Link Four</a>
                </div>
            </div>
            <div class="three wide column">
                <h4 class="ui inverted header">Group 2</h4>
                <div class="ui inverted link list">
                    <a href="#" class="item">Link One</a>
                    <a href="#" class="item">Link Two</a>
                    <a href="#" class="item">Link Three</a>
                    <a href="#" class="item">Link Four</a>
                </div>
            </div>
            <div class="three wide column">
                <h4 class="ui inverted header">Group 3</h4>
                <div class="ui inverted link list">
                    <a href="#" class="item">Link One</a>
                    <a href="#" class="item">Link Two</a>
                    <a href="#" class="item">Link Three</a>
                    <a href="#" class="item">Link Four</a>
                </div>
            </div>
            <div class="seven wide column">
                <h4 class="ui inverted header">Footer Header</h4>
                <p>Extra space for a call to action inside the footer that could help re-engage users.</p>
            </div>
        </div>
        <div class="ui inverted section divider"></div>
        <img src="http://semantic-ui.com/examples/assets/images/logo.png" class="ui centered mini image">
        <div class="ui horizontal inverted small divided link list">
            <a class="item" href="#">Site Map</a>
            <a class="item" href="#">Contact Us</a>
            <a class="item" href="#">Terms and Conditions</a>
            <a class="item" href="#">Privacy Policy</a>
        </div>
    </div>
</div>


</body>
</html>
