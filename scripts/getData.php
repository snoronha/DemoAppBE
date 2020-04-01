<?php

main();

function main() {

    $mysqli = new mysqli('localhost', 'root', 'root', 'DemoApp');
    /* check connection */
    if (mysqli_connect_errno()) {
        die("Connection failed: " . $mysqli->connect_error);
    }


    $urls = [
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=fruit&count=200',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=fruit&count=200&page=2',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=fruit&count=200&page=3',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=vegetables&count=200',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=vegetables&count=200&page=2',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=vegetables&count=200&page=3',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=vegetables&count=200&page=4',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=vegetables&count=200&page=5',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=household+essentials&count=200',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=household+essentials&count=200&page=2',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=household+essentials&count=200&page=3',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=meat&count=200',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=meat&count=200&page=2',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=meat&count=200&page=3',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=dairy&count=200',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=dairy&count=200&page=2',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=dairy&count=200&page=3',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=dairy&count=200&page=4',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=bread&count=200',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=pantry&count=200',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=pantry&count=200&page=2',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=snacks&count=200',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=snacks&count=200&page=2',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=snacks&count=200&page=3',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=snacks&count=200&page=4',
        'https://grocery.walmart.com/v4/api/products/search?storeId=2119&query=snacks&count=200&page=5',
    ];

    $allProds = [];

    for ($i = 0; $i < count($urls); $i++) {
        $url = $urls[$i];
        $flatProds = getData($url);
        for ($j = 0; $j < count($flatProds); $j++) {
            $product = $flatProds[$j];
            if (!isset($allProds[$product['USItemId']])) {
                $allProds[$product['USItemId']] = $product;
            }
        }
        printf("Count = %d\n", count($allProds));
        sleep(2);
    }
    // saveToCSV($allProds);
    saveToDB($mysqli, $allProds);
    $mysqli->close();
}

function saveToDB($mysqli, $products) {
    $prodKeys  = array_keys($products);
    $product   = $products[$prodKeys[0]];
    $headers =  [
        'USItemId', 'offerId', 'sku', 'salesUnit', 'name', 'name_lc',
        'thumbnail', 'weightIncrement', 'averageWeight', 'maxAllowed', 'productUrl',
        'isSnapEligible', 'type', 'rating', 'reviewsCount', 'isOutOfStock', 
        'list', 'previousPrice', 'priceUnitOfMeasure', 'salesUnitOfMeasure', 'salesQuantity',
        'displayCondition', 'displayPrice', 'displayUnitPrice', 'isClearance', 'isRollback', 'unit',
    ];
    $stmt = $mysqli->prepare(
        "INSERT INTO item (USItemId, offerId, sku, salesUnit, name, name_lc,
                           thumbnail, weightIncrement, averageWeight, maxAllowed, productUrl,
                           isSnapEligible, type, rating, reviewsCount, isOutOfStock,
                           list, previousPrice, priceUnitOfMeasure, salesUnitOfMeasure, salesQuantity,
                           displayCondition, displayPrice, displayUnitPrice, isClearance, isRollback, unit )
         VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
                 ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
    );

    foreach( $products as $p ) {
        foreach($headers as $h) {
            if ($p[$h] == "NULL") {
                $p[$h] = null;
            }
        }
        $stmt->bind_param('sssssssddisisdisddssisdssss', 
            $p['USItemId'], $p['offerId'], $p['sku'], $p['salesUnit'], $p['name'], $p['name_lc'],
            $p['thumbnail'], $p['weightIncrement'], $p['averageWeight'], $p['maxAllowed'], $p['productUrl'],
            $p['isSnapEligible'], $p['type'], $p['rating'], $p['reviewsCount'], $p['isOutOfStock'], 
            $p['list'], $p['previousPrice'], $p['priceUnitOfMeasure'], $p['salesUnitOfMeasure'], $p['salesQuantity'],
            $p['displayCondition'], $p['displayPrice'], $p['displayUnitPrice'], $p['isClearance'], 
            $p['isRollback'], $p['unit']
        );
        $stmt->execute();
    }
}

function saveToCSV($products) {
    $prodKeys  = array_keys($products);
    $product   = $products[$prodKeys[0]];

    $fileData  = [];
    $headers   = array_keys($product);
    $fileData[] = $headers;

    foreach( $products as $product ) {
        $rowData = [];
        for ( $j = 0; $j < count($headers); $j++ ) {
            $header = $headers[$j];
            $rowData[] = $product[$header];
        }
        $fileData[] = $rowData;
    }
    
    $fp = fopen('file.csv', 'w');
    foreach ($fileData as $row) {
        fputcsv($fp, $row);
    }
    fclose($fp);
}


function getData($url) {
    
    $ch = curl_init();
    $timeout = 5;

    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, $timeout);

    // Get URL content
    $page_string = curl_exec($ch);
    $page_json   = json_decode($page_string);
    curl_close($ch);

    $products  = $page_json->{'products'};
    $flatProds = [];
    for ($i = 0; $i < count($products); $i++) {
        $product  = $products[$i];
        $basic    = $product->{'basic'};
        $detailed = $product->{'detailed'};
        $store    = $product->{'store'};
        $price    = $store->{'price'};
        $obj = array(
            "USItemId"             => $product->{'USItemId'},
            "offerId"              => $product->{'offerId'},
            "sku"                  => isset($product->{'sku'}) ? $product->{'sku'} : "NULL",
            "salesUnit"            => $basic->{'salesUnit'} ? $basic->{'salesUnit'} : "NULL",
            "name"                 => $basic->{'name'},
            "name_lc"              => stemString(strtolower($basic->{'name'})),
            "thumbnail"            => $basic->{'image'}->{'thumbnail'},
            "weightIncrement"      => $basic->{'weightIncrement'} ? $basic->{'weightIncrement'} : "NULL",
            "averageWeight"        => $basic->{'averageWeight'} ? $basic->{'averageWeight'} : "NULL",
            "maxAllowed"           => $basic->{'maxAllowed'} ? $basic->{'maxAllowed'} : "NULL",
            "productUrl"           => $basic->{'productUrl'} ? $basic->{'productUrl'} : "NULL",
            "isSnapEligible"       => $basic->{'isSnapEligible'} ? $basic->{'isSnapEligible'} : "NULL",
            "type"                 => $basic->{'type'} ? $basic->{'type'} : "NULL",
            "rating"               => isset($detailed->{'rating'}) ? $detailed->{'rating'} : "NULL",
            "reviewsCount"         => isset($detailed->{'reviewsCount'}) ? $detailed->{'reviewsCount'} : "NULL",
            "isOutOfStock"         => $store->{'isOutOfStock'} ? $store->{'isOutOfStock'} : "NULL",
            "list"                 => $price->{'list'} ? $price->{'list'} : "NULL",
            "previousPrice"        => $price->{'previousPrice'} ? $price->{'previousPrice'} : "NULL",
            "priceUnitOfMeasure"   => $price->{'priceUnitOfMeasure'} ? $price->{'priceUnitOfMeasure'} : "NULL",
            "salesUnitOfMeasure"   => $price->{'salesUnitOfMeasure'} ? $price->{'salesUnitOfMeasure'} : "NULL",
            "salesQuantity"        => $price->{'salesQuantity'} ? $price->{'salesQuantity'} : "NULL",
            "displayCondition"     => $price->{'displayCondition'} ? $price->{'displayCondition'} : "NULL",
            "displayPrice"         => $price->{'displayPrice'} ? $price->{'displayPrice'} : "NULL",
            "displayUnitPrice"     => $price->{'displayUnitPrice'} ? $price->{'displayUnitPrice'} : "NULL",
            "isClearance"          => $price->{'isClearance'} ? $price->{'isClearance'} : "NULL",
            "isRollback"           => $price->{'isRollback'} ? $price->{'isRollback'} : "NULL",
            "unit"                 => $price->{'unit'} ? $price->{'unit'} : "NULL",
        );
        $flatProds[] = $obj;
    }
    return $flatProds;
}

function stemString($str) {
    $stemmedStr = "";
    $words = preg_split("/[\s,]+/", $str);
    $arr   = [];
    foreach($words as $word) {
        $stemmedWord = PorterStemmer::Stem($word);
        $arr[] = $stemmedWord;
    }
    return implode(" ", $arr);
}

/**
* Copyright (c) 2005 Richard Heyes (http://www.phpguru.org/)
*
* All rights reserved.
*
* This script is free software.
*/

/**
* PHP5 Implementation of the Porter Stemmer algorithm. Certain elements
* were borrowed from the (broken) implementation by Jon Abernathy.
*
* Usage:
*
*  $stem = PorterStemmer::Stem($word);
*
* How easy is that?
*/

class PorterStemmer {
    /**
    * Regex for matching a consonant
    * @var string
    */
    private static $regex_consonant = '(?:[bcdfghjklmnpqrstvwxz]|(?<=[aeiou])y|^y)';


    /**
    * Regex for matching a vowel
    * @var string
    */
    private static $regex_vowel = '(?:[aeiou]|(?<![aeiou])y)';


    /**
    * Stems a word. Simple huh?
    *
    * @param  string $word Word to stem
    * @return string       Stemmed word
    */
    public static function Stem($word)
    {
        if (strlen($word) <= 2) {
            return $word;
        }

        $word = self::step1ab($word);
        $word = self::step1c($word);
        $word = self::step2($word);
        $word = self::step3($word);
        $word = self::step4($word);
        $word = self::step5($word);

        return $word;
    }


    /**
    * Step 1
    */
    private static function step1ab($word)
    {
        // Part a
        if (substr($word, -1) == 's') {

                self::replace($word, 'sses', 'ss')
            OR self::replace($word, 'ies', 'i')
            OR self::replace($word, 'ss', 'ss')
            OR self::replace($word, 's', '');
        }

        // Part b
        if (substr($word, -2, 1) != 'e' OR !self::replace($word, 'eed', 'ee', 0)) { // First rule
            $v = self::$regex_vowel;

            // ing and ed
            if (   preg_match("#$v+#", substr($word, 0, -3)) && self::replace($word, 'ing', '')
                OR preg_match("#$v+#", substr($word, 0, -2)) && self::replace($word, 'ed', '')) { // Note use of && and OR, for precedence reasons

                // If one of above two test successful
                if (    !self::replace($word, 'at', 'ate')
                    AND !self::replace($word, 'bl', 'ble')
                    AND !self::replace($word, 'iz', 'ize')) {

                    // Double consonant ending
                    if (    self::doubleConsonant($word)
                        AND substr($word, -2) != 'll'
                        AND substr($word, -2) != 'ss'
                        AND substr($word, -2) != 'zz') {

                        $word = substr($word, 0, -1);

                    } else if (self::m($word) == 1 AND self::cvc($word)) {
                        $word .= 'e';
                    }
                }
            }
        }

        return $word;
    }


    /**
    * Step 1c
    *
    * @param string $word Word to stem
    */
    private static function step1c($word)
    {
        $v = self::$regex_vowel;

        if (substr($word, -1) == 'y' && preg_match("#$v+#", substr($word, 0, -1))) {
            self::replace($word, 'y', 'i');
        }

        return $word;
    }


    /**
    * Step 2
    *
    * @param string $word Word to stem
    */
    private static function step2($word)
    {
        switch (substr($word, -2, 1)) {
            case 'a':
                    self::replace($word, 'ational', 'ate', 0)
                OR self::replace($word, 'tional', 'tion', 0);
                break;

            case 'c':
                    self::replace($word, 'enci', 'ence', 0)
                OR self::replace($word, 'anci', 'ance', 0);
                break;

            case 'e':
                self::replace($word, 'izer', 'ize', 0);
                break;

            case 'g':
                self::replace($word, 'logi', 'log', 0);
                break;

            case 'l':
                    self::replace($word, 'entli', 'ent', 0)
                OR self::replace($word, 'ousli', 'ous', 0)
                OR self::replace($word, 'alli', 'al', 0)
                OR self::replace($word, 'bli', 'ble', 0)
                OR self::replace($word, 'eli', 'e', 0);
                break;

            case 'o':
                    self::replace($word, 'ization', 'ize', 0)
                OR self::replace($word, 'ation', 'ate', 0)
                OR self::replace($word, 'ator', 'ate', 0);
                break;

            case 's':
                    self::replace($word, 'iveness', 'ive', 0)
                OR self::replace($word, 'fulness', 'ful', 0)
                OR self::replace($word, 'ousness', 'ous', 0)
                OR self::replace($word, 'alism', 'al', 0);
                break;

            case 't':
                    self::replace($word, 'biliti', 'ble', 0)
                OR self::replace($word, 'aliti', 'al', 0)
                OR self::replace($word, 'iviti', 'ive', 0);
                break;
        }

        return $word;
    }


    /**
    * Step 3
    *
    * @param string $word String to stem
    */
    private static function step3($word)
    {
        switch (substr($word, -2, 1)) {
            case 'a':
                self::replace($word, 'ical', 'ic', 0);
                break;

            case 's':
                self::replace($word, 'ness', '', 0);
                break;

            case 't':
                    self::replace($word, 'icate', 'ic', 0)
                OR self::replace($word, 'iciti', 'ic', 0);
                break;

            case 'u':
                self::replace($word, 'ful', '', 0);
                break;

            case 'v':
                self::replace($word, 'ative', '', 0);
                break;

            case 'z':
                self::replace($word, 'alize', 'al', 0);
                break;
        }

        return $word;
    }


    /**
    * Step 4
    *
    * @param string $word Word to stem
    */
    private static function step4($word)
    {
        switch (substr($word, -2, 1)) {
            case 'a':
                self::replace($word, 'al', '', 1);
                break;

            case 'c':
                    self::replace($word, 'ance', '', 1)
                OR self::replace($word, 'ence', '', 1);
                break;

            case 'e':
                self::replace($word, 'er', '', 1);
                break;

            case 'i':
                self::replace($word, 'ic', '', 1);
                break;

            case 'l':
                    self::replace($word, 'able', '', 1)
                OR self::replace($word, 'ible', '', 1);
                break;

            case 'n':
                    self::replace($word, 'ant', '', 1)
                OR self::replace($word, 'ement', '', 1)
                OR self::replace($word, 'ment', '', 1)
                OR self::replace($word, 'ent', '', 1);
                break;

            case 'o':
                if (substr($word, -4) == 'tion' OR substr($word, -4) == 'sion') {
                    self::replace($word, 'ion', '', 1);
                } else {
                    self::replace($word, 'ou', '', 1);
                }
                break;

            case 's':
                self::replace($word, 'ism', '', 1);
                break;

            case 't':
                    self::replace($word, 'ate', '', 1)
                OR self::replace($word, 'iti', '', 1);
                break;

            case 'u':
                self::replace($word, 'ous', '', 1);
                break;

            case 'v':
                self::replace($word, 'ive', '', 1);
                break;

            case 'z':
                self::replace($word, 'ize', '', 1);
                break;
        }

        return $word;
    }


    /**
    * Step 5
    *
    * @param string $word Word to stem
    */
    private static function step5($word)
    {
        // Part a
        if (substr($word, -1) == 'e') {
            if (self::m(substr($word, 0, -1)) > 1) {
                self::replace($word, 'e', '');

            } else if (self::m(substr($word, 0, -1)) == 1) {

                if (!self::cvc(substr($word, 0, -1))) {
                    self::replace($word, 'e', '');
                }
            }
        }

        // Part b
        if (self::m($word) > 1 AND self::doubleConsonant($word) AND substr($word, -1) == 'l') {
            $word = substr($word, 0, -1);
        }

        return $word;
    }


    /**
    * Replaces the first string with the second, at the end of the string. If third
    * arg is given, then the preceding string must match that m count at least.
    *
    * @param  string $str   String to check
    * @param  string $check Ending to check for
    * @param  string $repl  Replacement string
    * @param  int    $m     Optional minimum number of m() to meet
    * @return bool          Whether the $check string was at the end
    *                       of the $str string. True does not necessarily mean
    *                       that it was replaced.
    */
    private static function replace(&$str, $check, $repl, $m = null)
    {
        $len = 0 - strlen($check);

        if (substr($str, $len) == $check) {
            $substr = substr($str, 0, $len);
            if (is_null($m) OR self::m($substr) > $m) {
                $str = $substr . $repl;
            }

            return true;
        }

        return false;
    }


    /**
    * What, you mean it's not obvious from the name?
    *
    * m() measures the number of consonant sequences in $str. if c is
    * a consonant sequence and v a vowel sequence, and <..> indicates arbitrary
    * presence,
    *
    * <c><v>       gives 0
    * <c>vc<v>     gives 1
    * <c>vcvc<v>   gives 2
    * <c>vcvcvc<v> gives 3
    *
    * @param  string $str The string to return the m count for
    * @return int         The m count
    */
    private static function m($str)
    {
        $c = self::$regex_consonant;
        $v = self::$regex_vowel;

        $str = preg_replace("#^$c+#", '', $str);
        $str = preg_replace("#$v+$#", '', $str);

        preg_match_all("#($v+$c+)#", $str, $matches);

        return count($matches[1]);
    }


    /**
    * Returns true/false as to whether the given string contains two
    * of the same consonant next to each other at the end of the string.
    *
    * @param  string $str String to check
    * @return bool        Result
    */
    private static function doubleConsonant($str)
    {
        $c = self::$regex_consonant;

        return preg_match("#$c{2}$#", $str, $matches) AND $matches[0][0] == $matches[0][1];
    }


    /**
    * Checks for ending CVC sequence where second C is not W, X or Y
    *
    * @param  string $str String to check
    * @return bool        Result
    */
    private static function cvc($str)
    {
        $c = self::$regex_consonant;
        $v = self::$regex_vowel;

        return     preg_match("#($c$v$c)$#", $str, $matches)
                AND strlen($matches[1]) == 3
                AND $matches[1][2] != 'w'
                AND $matches[1][2] != 'x'
                AND $matches[1][2] != 'y';
    }
}

?>