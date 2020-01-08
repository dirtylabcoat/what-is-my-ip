# what-is-my-ip
Cloud function that returns the caller's IP-address.

Deploy with:

	gcloud functions deploy whatismyip --runtime go111 --allow-unauthenticated --trigger-http \
	--entry-point WhatIsMyIp --memory 128MB --region europe-west2

Test with:

	curl https://REGION-PROJECT.cloudfunctions.net/whatismyip

