<template>
    <div class="home">
        <div id="one"></div>

        <div class="two">
            <div id="twoLeft"></div>
            <div id="twoRight"></div>
        </div>
        <div class="three">
            <div id="threeLeft"></div>
            <div id="threeCenter"></div>
            <div id="threeRight"></div>
        </div>
      
    </div>
</template>
<script>
export default {
    data() {
        return {
            one: {
                title: {
                    text: '世界人口总量',
                    subtext: '数据来自网络'
                },
                tooltip: {
                    trigger: 'axis',
                    axisPointer: {
                        type: 'shadow'
                    }
                },
                legend: {
                    data: ['2011年', '2012年']
                },
                grid: {
                    left: '3%',
                    right: '4%',
                    bottom: '3%',
                    containLabel: true
                },
                xAxis: {
                    type: 'value',
                    boundaryGap: [0, 0.01]
                },
                yAxis: {
                    type: 'category',
                    data: ['巴西', '印尼', '美国', '印度', '中国', '世界人口(万)']
                },
                series: [
                    {
                        name: '2011年',
                        type: 'bar',
                        data: [18203, 23489, 29034, 104970, 131744, 330230]
                    },
                    {
                        name: '2012年',
                        type: 'bar',
                        data: [19325, 23438, 31000, 121594, 134141, 381807]
                    }
                ]
            },

            twoLeft: {
                xAxis: {
                    type: 'category',
                    boundaryGap: false,
                    data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
                },
                yAxis: {
                    type: 'value'
                },
                series: [{
                    data: [820, 932, 901, 934, 1290, 1330, 1320],
                    type: 'line',
                    areaStyle: {}
                }]
            },
            twoRight: {
                tooltip: {
                    formatter: '{a} <br/>{b} : {c}%'
                },
                series: [{
                    name: 'Pressure',
                    type: 'gauge',
                    detail: {
                        formatter: '{value}'
                    },
                    data: [{
                        value: 50,
                        name: 'SCORE'
                    }]
                }]
            },
        }
    },

    mounted(){
        this.oneEcharts()
    },
    methods: {
        //this.$root => app
        oneEcharts() {
            let one = document.getElementById("one")
            let oneChart = this.$echarts.init(one);
            oneChart.setOption(this.one);

            let twoLeft = document.getElementById("twoLeft")
            let twoLeftChart = this.$echarts.init(twoLeft);
            twoLeftChart.setOption(this.twoLeft);

            let twoRight = document.getElementById("twoRight")
            let twoRightChart = this.$echarts.init(twoRight, 'dark');
            twoRightChart.setOption(this.twoRight);
        },
        
    },
}
</script>

<style scoped>
.home {
    height: 100%;
}

.home #one{
    height: 30%;
    background-color: #fff;
    border-radius: 5px;
    margin: 10px 10px;
}

.home .two,
.home .three {
    height: 30%;
    display: flex;
    justify-content: space-between;
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