<?php
header("Content-type:text/html;charset=utf-8");
$param['name'] = 'mr.chen';
$param['pwd']  = '123456';






echo 'sign: '.createSign($param);


/**
 * 生成签名
 * @param array $params
 * @return string
 */
function createSign(Array $params){
    // 将删除参数组中所有等值为FALSE的参数（包括：NULL, 空字符串，0， false）
    $params = array_filter($params);
    // 按照键名对参数数组进行升序排序
    ksort($params);      
    //$str = http_build_query($params);  

    $str = urldecode(http_build_query($params));  
  	$secret = 'f7rbsEwiO5bl1aTd';
	$sign = md5(md5($str).md5($secret));  
    return $sign;
}