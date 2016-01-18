<style>
body { 
    padding-top: 70px; 
}

li a.active {
    font-weight: 900;
    color: #563d7c !important;
}

.navbar-collapse.collapse {
    display: block !important;
}

.navbar-nav > li, .navbar-nav {
    float: left !important;
}

.navbar-nav.navbar-right:last-child {
    margin-right: -15px !important;
}

.navbar-right {
    float: right !important;
}
</style>

<template>
<div id="wrapper" class="container">
    <nav class="navbar navbar-default navbar-fixed-top">
        <div class="container-fluid">
            <ul class="nav navbar-nav">
                <li><a v-link="'/posts'">Posts</a></li>
            </ul>
        </div>
    </nav>
    
    <router-view
     keep-alive
     transition-mode="out-in">
    </router-view>
</div>
</template>

