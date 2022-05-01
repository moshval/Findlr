<script>
    let newItem = '';
	let glet = Array(5);
	let ylet = Array(5);
	let xlet = '';
	let wordList = [];
	let wordList2 = [];
	let lan = 0;
	
	function addToList() {
		if(wordList == null || wordList.length == 0) {
			wordList2 = [];
			return
		}
		wordList2 = []
		console.log("called")
		for (let i = 0; i < wordList.length; i++) {
			let obj = {text: wordList[i], status: false};
			wordList2.push(obj);
		}
		console.log(wordList2);
	}
	function changeLang(){
		let lg = document.getElementById("lang")
		console.log(lg.value)
		if(lg.value == "en"){
			lan = 0;
		}
		else if(lg.value == "id"){
			lan = 1;
		}
	}
	function removeFromList(index) {
		wordList2.splice(index, 1)
		wordList2 = wordList2;
    }
	async function fetchWords(lang,green,yellow,exc){
		let gr = [];
		let yl = [];
		for(let i=0;i<5;i++){
			gr[i] = green[i];
			yl[i] = yellow[i];
			if(green[i] == '' || green[i] == null){
				gr[i] = '_';
			}
			if(yellow[i] == '' || yellow[i] == null){
				yl[i] = '_';
			}
		}
		let grn = gr.join('');
		console.log(grn);
		let ylw = yl.join('');
		console.log(ylw);
		grn = grn.toLowerCase();
		ylw = ylw.toLowerCase();
		exc = exc.replace(/\s/g, "")
		let url;
		if(exc.length > 0){
			url = 'https://findlr-bend.herokuapp.com/words?'+'type='+lang+'&word='+grn+'&bad='+ylw+'&exc='+exc;
		}
		else{
			url = 'https://findlr-bend.herokuapp.com/words?'+'type='+lang+'&word='+grn+'&bad='+ylw;
		}
		console.log(url)
		let response = await fetch(url)
    		.then(response => response.json())  // convert to json
    		.then(json => wordList = json.words)    //print data to console
    		.catch(err => console.log('Request Failed', err)); // Catch error
		addToList();
	}
</script>
<div class="container">
	<div class="container2">
	<div class="sub-container">
	<h1 style="text-align:center">Findlr</h1>
	<h2>Lang</h2>
	<select id="lang" on:change={() => changeLang()}>
		<option value="en">English</option>
		<option value="id">Indonesian</option>
	</select>
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
	<button class="sbmt-btn" on:click={fetchWords(lan,glet,ylet,xlet)}>Submit</button>
	<h2>Possible Answer(s)</h2>
	<div class ="answers">
	<ul>
	{#each wordList2 as item, index}
		<li>
		<input bind:checked={item.status} type="checkbox">
		<span class:checked={item.status}>{item.text}</span>
		<span id="del-item" on:click={() => removeFromList(index)}>❌</span>
		</li>
	{/each}
	</ul>	
	</div>
	</div>
	</div>
</div>


<style> 
	.checked {
        text-decoration: line-through;
    }
	.green-letters{
		display: flex;
	}
	.green-letters .colored-letter{
		background-color: #40a375;
		color: white;
	}
	.yellow-letters .colored-letter{
		background-color: #ffba01;
		color: black;
	}
	.yellow-letters{
		display: flex;
	}
	.excluded-letter{
		width: 16rem;
	}
	ul{
		list-style: none;
		display: block;
		padding-left: 0;
		padding-top: 0;
		-moz-column-count: 2;
		-moz-column-gap: 10px;
		-webkit-column-count: 2;
		-webkit-column-gap: 10px;
		column-count: 2;
		column-gap: 10px;
	}
	.sbmt-btn{
		width: 16rem;
	}
	.sbmt-btn:hover{
		cursor: pointer;
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
	.container2{
		border-radius: 20px;
	}
	.sub-container{
		align-items: center;
		padding: 4rem;
	}
	#del-item{
		cursor: pointer;
	}

</style> 