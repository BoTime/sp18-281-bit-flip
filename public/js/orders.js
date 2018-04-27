    //To call delete of a user order
	function CallDelete(pid){
		var path = "http://"+$(location).attr('host')+"/orders";
		console.log("Delete API ", path);
		dataJSON = JSON.stringify({"pid": pid});
		$.ajax({
			data: dataJSON,
			dataType: "json",
						type: "DELETE",
			url: path,
			success: function(data){
				alert("Order Deleted");
				window.location.href=window.location.href;
			},
			error: function(){
				alert("Something went wrong, try again");
			}
		});
	};	
	//Dynamically create rows for a user orders to view
	function CreateTableFromJSON() {
		var myOrders  ;
		var path = "http://"+$(location).attr('host')+"/orders";
		console.log(path);
		$.ajax({type: "GET", url: path, success: function(result){
			myOrders =result;
			console.log(myOrders)
		// ADD JSON DATA TO THE DIV AS ROWS.
			for (var i = 0; myOrders != null && i < myOrders.length; i++) {
				 // CREATE DYNAMIC DIV.
				var brk = document.createElement("br"); 			
				var table = document.createElement("div");
				var rowid = "ul"+i;
				var list = document.createElement("ol");
				var line1 = document.createElement("h3");
				if(myOrders[i].store === 'store1'){
					line1.innerHTML  = "<span style='color: red;text-transform:capitalize'>"+myOrders[i].status+"</span> at San Jose store";
				}else{
					line1.innerHTML  = "<span style='color: red;text-transform:capitalize'>"+myOrders[i].status+"</span> at Mountain View store";
				}
				var button = document.createElement("button");
				var span = document.createElement('span');
				span.innerHTML = '<button id="' + myOrders[i].pay_id +'" onclick="CallDelete(this.id)" />Delete';
				
				//If orders status is placed, not allow user to delete them
				if(myOrders[i].status === 'placed'){
						span.innerHTML = '';
				}
				//To list items elected in an order
				var line2 = document.createElement("p");
				line2.innerHTML = "Items:";
				for (var k = 0; k < myOrders[i].product.length; k++) {
						var x = document.createElement("LI");
						console.log(myOrders[i].product[k].item);
						var lists = myOrders[i].product[k].item + " , " + myOrders[i].product[k].qty+" , "+ myOrders[i].product[k].size;
						var t = document.createTextNode(lists);
						x.appendChild(t);
						list.append(x);
				}
				table.append(line1);
				table.append(span);
				table.append(line2);
				table.append(list);
				table.append(brk);
				// FINALLY ADD THE NEWLY CREATED TABLE WITH JSON DATA TO A CONTAINER.
				$("#showData").append(table);
				console.log("Appended", i, table);
			}        
		},
		complete: function(xhr, textStatus) {
			console.log(xhr.status);
			if (xhr.status === 400 || xhr.status === 401 || xhr.status === 404 ){
				window.location="http://"+$(location).attr('host')+"/signin";
			}else{
				window.location="http://"+$(location).attr('host')+"/oops";
			}	
		} 
	});
		
    }
