var nav_bar = {
    template: '<nav class="navbar">\n' +
        '      <button class="btn btn-default navbar-btn fa fa-bars"></button>\n' +
        '      <ul class="nav navbar-nav navbar-right">\n' +
        '        <li><a href="/admin/profile.html"><i class="fa fa-user"></i>个人中心</a></li>\n' +
        '        <li><a href="/admin/login/loginOut"><i class="fa fa-sign-out"></i>退出</a></li>\n' +
        '      </ul>\n' +
        '    </nav>',
    data:function () {
        return {}
    }
};

var head_image = {
    template:'<div class="profile">\n' +
        '        <img class="avatar" v-bind:src="avatar">\n' +
        '        <h3 class="name">{{ name }}</h3>\n' +
        '      </div>',
    data:function () {
        return {avatar:'/uploads/avatar.jpg', name:''}
    },
    methods:{
        getUserInfo(){ //获取用户信息判断是否已登录
            let userUrl = '/admin/user/index';
            this.$http.get(userUrl).then(response => {
                if (response.data.code == 0) {
                    let userInfo = response.data.data
                    this.avatar = userInfo.head_image ? userInfo.head_image : '/uploads/avatar.jpg'
                        this.name = userInfo.email ? userInfo.email : (userInfo.mobile ? userInfo.mobile : 'dfh420984')
                } else {
                    window.location.href='/admin/login.html'
                }
            }, response => {
                // error callback
                if (response.data.code == 403) {
                    window.location.href='/admin/login.html'
                }
            });
        }
    },
    created(){
        this.getUserInfo()
    }
}

var aside_temple = {
    template:'<ul class="nav">\n' +
        '      <li>\n' +
        '      <a href="/admin/index.html"><i class="fa fa-dashboard"></i>仪表盘</a>\n' +
        '      </li>\n' +
        '      <li class="active">\n' +
        '      <a href="#menu-posts" data-toggle="collapse">\n' +
        '      <i class="fa fa-thumb-tack"></i>文章<i class="fa fa-angle-right"></i>\n' +
        '      </a>\n' +
        '      <ul id="menu-posts" class="collapse in">\n' +
        '      <li class="active"><a href="/admin/posts.html">所有文章</a></li>\n' +
        '      <li><a href="/admin/post-add.html">写文章</a></li>\n' +
        '      <li><a href="/admin/categories.html">分类目录</a></li>\n' +
        '      </ul>\n' +
        '      </li>\n' +
        '      <li>\n' +
        '      <a href="/admin/comments.html"><i class="fa fa-comments"></i>评论</a>\n' +
        '      </li>\n' +
        '      <li>\n' +
        '      <a href="/admin/users.html"><i class="fa fa-users"></i>用户</a>\n' +
        '      </li>\n' +
        '      <li>\n' +
        '      <a href="#menu-settings" class="collapsed" data-toggle="collapse">\n' +
        '      <i class="fa fa-cogs"></i>设置<i class="fa fa-angle-right"></i>\n' +
        '      </a>\n' +
        '      <ul id="menu-settings" class="collapse">\n' +
        '      <li><a href="/admin/nav-menus.html">导航菜单</a></li>\n' +
        '      <li><a href="/admin/slides.html">图片轮播</a></li>\n' +
        '      <li><a href="/admin/settings.html">网站设置</a></li>\n' +
        '      </ul>\n' +
        '      </li>\n' +
        '      </ul>'
}
Vue.component('nav_bar',nav_bar);
Vue.component('head_image',head_image);
Vue.component('aside_temple',aside_temple);