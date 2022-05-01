<script>
    let newItem = '';
	let glet = Array(5);
	let ylet = Array(5);
	let xlet = '';
	
    let todoList = [{text: 'Write my first post', status: true},
                    {text: 'Upload the post to the blog', status: false},
                    {text: 'Publish the post at Facebook', status: false}];
	
	function addToList() {
		todoList = [...todoList, {text: newItem, status: false}];
		newItem = '';
	}
	
	function removeFromList(index) {
		todoList.splice(index, 1)
		todoList = todoList;
    }
	async function fetchWords(lang,green,yellow,exc){
		for(let i=0;i<5;i++){
			if(green[i] == ''){
				green[i] = '_';
			}
			if(yellow[i] == ''){
				yellow[i] = '_';
			}
		}
		let grn = green.join('');
		console.log(grn);
		let ylw = yellow.join('');
		console.log(ylw);
		exc = exc.replace(/\s/g, "")
		let url;
		if(exc.length > 0){
			url = 'https://findlr-bend.herokuapp.com/words?'+'type='+lang+'&word='+grn+'&bad='+ylw+'&exc='+exc;
		}
		else{
			url = 'https://findlr-bend.herokuapp.com/words?'+'type='+lang+'&word='+grn+'&bad='+ylw;
		}
		let response = await fetch(url,{
			method:'GET',headers:{
				'Content-Type': 'application/json'
			}
		});
		console.log(response.json())
		return response.json();
	}
</script>
<div class="container">
	<div class="sub-container">
	<h1 style="text-align:center">Findlr</h1>
	<h2>Green Letters</h2>
	<div class="green-letters">
		<input class="colored-letter" bind:value={glet[0]} type="text" maxlength="1" placeholder="">
		<input class="colored-letter" bind:value={glet[1]} type="text" maxlength="1" placeholder="">
		<input class="colored-letter" bind:value={glet[2]} type="text" maxlength="1" placeholder="">
		<input class="colored-letter" bind:value={glet[3]} type="text" maxlength="1" placeholder="">
		<input class="colored-letter" bind:value={glet[4]} type="text" maxlength="1" placeholder="">
	</div>
	<h2>Yellow Letters</h2>
	<div class="yellow-letters">
		<input class="colored-letter" bind:value={ylet[0]} type="text" maxlength="1" placeholder="">
		<input class="colored-letter" bind:value={ylet[1]} type="text" maxlength="1" placeholder="">
		<input class="colored-letter" bind:value={ylet[2]} type="text" maxlength="1" placeholder="">
		<input class="colored-letter" bind:value={ylet[3]} type="text" maxlength="1" placeholder="">
		<input class="colored-letter" bind:value={ylet[4]} type="text" maxlength="1" placeholder="">
	</div>
	<h2>Excluded Letters</h2>
		<input class="excluded-letter" bind:value={xlet} type="text" placeholder="">
	<br>
	<button on:click={fetchWords(0,glet,ylet,xlet)}>Submit</button>

	<br/>
	{#each todoList as item, index}
		<input bind:checked={item.status} type="checkbox">
		<span class:checked={item.status}>{item.text}</span>
		<span on:click={() => removeFromList(index)}>❌</span>
		<br/>
	{/each} 
	</div>
</div>


<style> 
	.checked {
        text-decoration: line-through;
    }
	.green-letters{
		display: flex;
	}
	.yellow-letters{
		display: flex;
	}
	.excluded-letter{
		width: 16rem;
	}
	.colored-letter{
		border: 2px solid gray;
		border-radius: 3px;
		margin: 2px;
		font-size: 2.5rem;
		font-weight: 700;
		height: 3rem;
		width: 3rem;
		display: flex;
		padding: 0;
		text-align: center;
		justify-content: center;
		align-items: center;
		text-transform: uppercase;
	}
	.container{
		margin: auto;
		display:flex;
		align-items: center;
		justify-content: center;
	}
</style> 