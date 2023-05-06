<script setup>
import HelloWorld from './components/HelloWorld.vue'
import ConnectionList from './components/ConnectionList.vue'

</script>


<template>
	<el-row>
		<el-col :span="6" style="height: 100vh; background-color:rgb(95, 78, 78);">
			<ConnectionList/>
		</el-col>
	</el-row>
	<HelloWorld/>
</template>


