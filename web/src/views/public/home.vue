<template>
    <div class="home">
        <div id="one"> 
            等待补充....
        </div>

        <div class="two">
            <div id="twoLeft">等待补充....</div>
            <div id="twoRight">等待补充....</div>
        </div>
        <div class="three">
            <div id="threeLeft">等待补充....</div>
            <div id="threeCenter">等待补充....</div>
            <div id="threeRight">等待补充....</div>
        </div>
      
    </div>
</template>

<script>
export default {
    data() {
       return {}
    },
}
</script>

<style scoped>
.home {
    height: 655px;
    /* background: pink; */
}

.home #one{
    height: 31%;
    background-color: lightpink;
    border-radius: 5px;
    margin: 10px 10px;
}

.home .two,
.home .three {
    height: 30%;
    display: flex;
    justify-content: space-between;
    background-color: pink;
}

.home .two > div,
.home .three div {
    height: 90%;
    width: 50%;
    border-radius: 5px; 
    background-color: #fff; 
    margin: 10px 10px;
}
</style>