DEADLINKS
=========
Simple tool to look for 404s on a web page

USAGE
=====
./deadlinks [-a] <uri>
  -a return all status codes

CAVEATs
=======
For a large number of links you may want to increase the ulimit -n 15000


EXAMPLE
========
./deadlinks -a http://indiajoze.com/index.shtml
fetching...  http://indiajoze.com/questionnaire.html
fetching...  http://indiajoze.com/extended.html
fetching...  http://indiajoze.com/earlyadopters.html
fetching...  http://indiajoze.com/equipment.html
fetching...  http://indiajoze.com/joze.html
fetching...  http://indiajoze.com/chefs.html
fetching...  http://indiajoze.com/press.html
fetching...  http://indiajoze.com/order/group.html
fetching...  http://indiajoze.com/order/eat/
fetching...  http://indiajoze.com/blog.html
fetching...  http://indiajoze.com/festivals.shtml
fetching...  http://www.facebook.com/IndiaJoze
fetching...  http://indiajoze.com/catering.html
fetching...  http://indiajoze.com/recipes.html
fetching...  http://indiajoze.com/upcoming.html
[200]	http://indiajoze.com/upcoming.html	200 OK
[200]	http://indiajoze.com/order/eat/	200 OK
[200]	http://indiajoze.com/blog.html	200 OK
[200]	http://indiajoze.com/order/group.html	200 OK
[200]	http://indiajoze.com/questionnaire.html	200 OK
[200]	http://indiajoze.com/festivals.shtml	200 OK
[200]	http://indiajoze.com/recipes.html	200 OK
[200]	http://indiajoze.com/chefs.html	200 OK
[200]	http://indiajoze.com/extended.html	200 OK
[200]	http://indiajoze.com/catering.html	200 OK
[200]	http://indiajoze.com/equipment.html	200 OK
[200]	http://indiajoze.com/earlyadopters.html	200 OK
[200]	http://indiajoze.com/joze.html	200 OK
[200]	http://indiajoze.com/press.html	200 OK
[200]	http://www.facebook.com/IndiaJoze	200 OK
OK
