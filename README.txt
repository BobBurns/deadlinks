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
./deadlinks -a unr.edu
metilda deadlinks # ./deadlinks -a unr.edu
fetching...  https://www.bcnpurchasing.nevada.edu
fetching...  https://unr.edu/around-campus
fetching...  https://www.unr.edu/
fetching...  https://unr.edu/equal-opportunity-title-ix
fetching...  http://giving.unr.edu/
fetching...  https://unr.edu/
fetching...  https://unr.edu/general-information/site-map
fetching...  https://unr.edu
fetching...  https://www.unr.edu/images/unr-main/homepage-features/2018-Nevada-Photo-Social.jpg
fetching...  https://www.unr.edu/directory
fetching...  https://my.nevada.unr.edu
fetching...  https://www.unr.edu/images/unr-main/homepage-features/2018-Nevada-Wallpaper.jpg
fetching...  https://unr.edu/diversity
fetching...  https://webcampus.unr.edu/
fetching...  http://www.unr.edu/degrees
fetching...  https://unr.edu/grad
fetching...  https://unr.edu/student-life/living-on-campus
fetching...  https://oit.unr.edu/services-and-support/email-phone-and-communications/email/
fetching...  http://www.nevadawolfpack.com
fetching...  https://unr.edu/academic-central/academic-resources/schedules-and-catalogs
fetching...  https://library.unr.edu/
fetching...  http://www.youtube.com/playlist?list=PL2V4CDJPEZiF_raq0x_oKzmFmaMkiNWZ0
fetching...  https://www.unr.edu/faculty-staff-forms
fetching...  https://unr.edu/general-information/contact-us
fetching...  https://unr.edu/diversity/campus-resources
fetching...  https://unr.edu/about/visit-campus
fetching...  https://unr.edu/workday
fetching...  https://unr.edu/admissions/how-to-apply
fetching...  https://unr.edu/about
fetching...  https://unr.edu/about/tour
fetching...  http://www.formstack.com/forms/?1226969-V0nnGE7i8j&URL=https://www.unr.edu/
fetching...  https://unr.edu/admissions
fetching...  https://unr.edu/academics
fetching...  mailto:admissions@unr.edu
fetching...  https://unr.edu/honors
Errors found:  Get mailto:admissions@unr.edu: unsupported protocol scheme "mailto"
fetching...  https://www.facebook.com/UniversityofNevada
fetching...  https://unr.edu/student-life
fetching...  https://unr.edu/research
fetching...  https://twitter.com/unevadareno
fetching...  https://unr.edu/impact
fetching...  https://unr.edu/nevada-today
fetching...  https://events.unr.edu/
fetching...  http://www.youtube.com/user/universityofnevada
fetching...  http://www.flickr.com/photos/unrphotos
fetching...  http://www.instagram.com/unevadareno
fetching...  https://www.linkedin.com/edu/university-of-nevada-reno-18908
fetching...  https://unr.edu/hr/jobs-and-working-at-nevada/job-opportunities
fetching...  https://unr.edu/general-information/emergency
fetching...  https://unr.edu/general-information/privacy
fetching...  https://unr.edu/accessibility/commitment
fetching...  https://unr.edu/general-information/copyright
[0]	mailto:admissions@unr.edu	
[200]	https://www.unr.edu/images/unr-main/homepage-features/2018-Nevada-Wallpaper.jpg	200 OK
[200]	https://www.unr.edu/faculty-staff-forms	200 OK
[200]	https://www.unr.edu/images/unr-main/homepage-features/2018-Nevada-Photo-Social.jpg	200 OK
[200]	https://www.bcnpurchasing.nevada.edu	200 OK
[200]	https://library.unr.edu/	200 OK
[200]	https://unr.edu/about	200 OK
[200]	http://www.unr.edu/degrees	200 OK
[200]	https://unr.edu/workday	200 OK
[200]	https://unr.edu/impact	200 OK
[200]	https://unr.edu/admissions	200 OK
[200]	https://unr.edu/grad	200 OK
[200]	https://www.unr.edu/directory	200 OK
[200]	https://unr.edu/general-information/privacy	200 OK
[200]	https://unr.edu/academics	200 OK
[200]	https://unr.edu/student-life/living-on-campus	200 OK
[200]	https://unr.edu/academic-central/academic-resources/schedules-and-catalogs	200 OK
[200]	https://unr.edu/general-information/copyright	200 OK
[200]	https://unr.edu/general-information/emergency	200 OK
[200]	https://unr.edu/hr/jobs-and-working-at-nevada/job-opportunities	200 OK
[200]	https://unr.edu/admissions/how-to-apply	200 OK
[200]	https://unr.edu/general-information/contact-us	200 OK
[200]	https://unr.edu/student-life	200 OK
[200]	https://my.nevada.unr.edu	200 OK
[200]	https://unr.edu/equal-opportunity-title-ix	200 OK
[200]	https://unr.edu/accessibility/commitment	200 OK
[200]	https://twitter.com/unevadareno	200 OK
[200]	https://unr.edu/honors	200 OK
[200]	http://www.nevadawolfpack.com	200 OK
[200]	https://events.unr.edu/	200 OK
[200]	https://oit.unr.edu/services-and-support/email-phone-and-communications/email/	200 OK
[200]	https://www.unr.edu/	200 OK
[200]	https://unr.edu/research	200 OK
[200]	https://unr.edu/	200 OK
[200]	https://unr.edu	200 OK
[200]	https://webcampus.unr.edu/	200 OK
[200]	http://giving.unr.edu/	200 OK
[200]	http://www.formstack.com/forms/?1226969-V0nnGE7i8j&URL=https://www.unr.edu/	200 OK
[200]	https://unr.edu/about/tour	200 OK
[200]	https://unr.edu/nevada-today	200 OK
[200]	https://unr.edu/diversity/campus-resources	200 OK
[200]	https://unr.edu/general-information/site-map	200 OK
[200]	http://www.youtube.com/playlist?list=PL2V4CDJPEZiF_raq0x_oKzmFmaMkiNWZ0	200 OK
[200]	https://unr.edu/about/visit-campus	200 OK
[200]	https://unr.edu/around-campus	200 OK
[200]	https://unr.edu/diversity	200 OK
[200]	http://www.instagram.com/unevadareno	200 OK
[200]	http://www.youtube.com/user/universityofnevada	200 OK
[999]	https://www.linkedin.com/edu/university-of-nevada-reno-18908	999 Request denied
[200]	https://www.facebook.com/UniversityofNevada	200 OK
[200]	http://www.flickr.com/photos/unrphotos	200 OK
OK
