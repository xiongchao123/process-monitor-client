<?php

while(true){
	file_put_contents("/tmp/Go/monitor/test/test.log",date("Y-m-d H:i:s").PHP_EOL,FILE_APPEND);
	sleep(10);
}
