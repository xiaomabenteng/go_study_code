<?php
$fp = stream_socket_client("tcp://127.0.0.1:8082",$errno, $errstr, 3);//连接 go的rpc server，其实就是socket连接
if (!$fp) {
    echo "$errstr ($errno)<br />\n";
    return;
}
 fwrite($fp, json_encode([
        'method' =>"UserService.GetUserName",
        'params' =>[101],
        'id'     => 0,
    ])."\n");

    echo fgets($fp);

fclose($fp);