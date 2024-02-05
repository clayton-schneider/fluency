window.addEventListener("load", () => {
	document.body.addEventListener("htmx:load", () => {
		if (document.querySelectorAll("tr").length) {
			buildSSChart(getAccelVals())
		}
	})
})


/**
	* Finds all acceleration values on current page
	* @constructor
	* @param {string} searchId
	*/
function getAccelVals() {
	// const parent = document.getElementById(searchId)
	const colEls = document.querySelectorAll("td:nth-child(1)")
	const colVals = []
	colEls.forEach(el => colVals.push(el.innerText))
	return colVals
}


function buildSSChart(measures) {
	const width = 1500;
	const height = 600;
	const marginTop = 20;
	const marginRight = 20;
	const marginBottom = 50;
	const marginLeft = 100;

	// Create the SVG container.
	const svg = d3.create("svg")
		.attr("id", "ssc-chart")
		.attr("width", width)
		.attr("height", height);

	// Declare the x (horizontal position) scale.
	const x = d3.scaleLinear().domain([0, 140]).range([marginLeft, width - marginRight])
	// Add the vertical lines on the graph
	for (let i = 1; i <= 140; i++) {
		svg.append("line")
			.attr("x1", x(i))
			.attr("y1", marginTop)
			.attr("x2", x(i))
			.attr("y2", height - marginBottom)
			.attr("stroke-width", i % 7 === 0 ? 0.5 : 0.2)
	}

	// declare the top-x scale
	const topX = d3.scaleLinear().domain([0, 20]).range([marginLeft, width - marginRight])

	// Declare the y (vertical position) scale.
	const y = d3.scaleLog().domain([0.0001, 1000]).range([height - marginBottom, marginTop])
	y.ticks().forEach(tick =>
		svg.append("line")
			.attr("x1", marginLeft)
			.attr("y1", y(tick))
			.attr("x2", width - marginRight)
			.attr("y2", y(tick))
			.attr("stroke-width", 0.2)
	)

	// Add the x-axis.
	svg.append("g")
		.attr("transform", `translate(0,${height - marginBottom})`)
		.call(d3.axisBottom(x));

	// Add bottom x-axis label
	svg.append("text")
		.attr("x", (width + marginLeft) / 2)
		.attr("y", height - 2)
		.attr("text-anchor", "middle")
		.text("SUCCESSIVE CALENDAR DAYS")

	// Add the top x-axis.
	svg.append("g")
		.attr("transform", `translate(0,${marginTop})`)
		.call(d3.axisTop(topX).tickValues([0, 4, 8, 12, 16, 20]));

	// Add the y-axis.
	svg.append("g")
		.attr("transform", `translate(${marginLeft},0)`)
		.call(d3.axisLeft(y).tickFormat(function(d) {
			return y.tickFormat(10, d3.format(""))(d)
		}))

	// Add left y-axis label
	svg.append("text")
		.attr("x", -(height / 2))
		.attr("y", marginLeft / 2)
		.attr("text-anchor", "middle")
		.attr("transform", "rotate(-90)")
		.text("COUNT PER MINUTE")


	measures.forEach((d, i) => {
		svg.append("circle")
			.attr("fill", "white")
			.attr("cx", x(i + 1))
			.attr("cy", y(d))
			.attr("r", 1);
	})




	// set colors
	const color = "rgb(68, 187, 239)"
	svg.style('stroke', color).style('color', color)

	// Append the SVG element.
	if (container.firstChild) {
		const curSVGEl = container.querySelector("svg")
		container.removeChild(curSVGEl)
		container.append(svg.node())
	} else {
		container.append(svg.node());
	}
}
