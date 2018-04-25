    
	function CallDelete(pid){
		var path = "http://"+$(location).attr('host')+"/orders";
		console.log(path);
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
	function CreateTableFromJSON() {
		var myOrders  ;
		var path = "http://"+$(location).attr('host')+"/orders";
		console.log(path);
		$.ajax({type: "GET", url: path, success: function(result){
			myOrders =result;
			console.log(myOrders)
		// ADD JSON DATA TO THE TABLE AS ROWS.
        for (var i = 0; myOrders != null && i < myOrders.length; i++) {
			 // CREATE DYNAMIC TABLE.
			var brk = document.createElement("br"); 
			var button = document.createElement("button");
			var span = document.createElement('span');
			span.innerHTML = '<button id="' + myOrders[i].pay_id +'" onclick="CallDelete(this.id)" />Delete';
			
			var table = document.createElement("div");
			var rowid = "ul"+i;
			var list = document.createElement("ol");
			var line1 = document.createElement("h3");
			if(myOrders[i].store === 'store1'){
				line1.innerHTML  = "<span style='color: red;text-transform:capitalize'>"+myOrders[i].status+"</span> at San Jose store";
			}else{
				line1.innerHTML  = "<span style='color: red;text-transform:capitalize'>"+myOrders[i].status+"</span> at Mountain View store";
			}

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

        
		}});
		
    }
