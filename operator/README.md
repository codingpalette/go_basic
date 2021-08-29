## 산술 연산자




<table>
    <thead>
        <tr>
            <th>구분</th>
            <th>연산자</th>
            <th>연산</th>
            <th>피연산자 타입</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td rowspan=5>사칙 </br> 연산과 </br> 나머지</td>
            <td>+</td>
            <td>덧셈</td>
            <td>정수,필수,복소수,문자열</td>
        </tr>
        <tr>
            <td>-</td>
            <td>뺄셈</td>
            <td>정수,실수,복소수</td>
        </tr>
        <tr>
            <td>*</td>
            <td>곱셈</td>
            <td>정수,실수,복소수</td>
        </tr>
        <tr>
            <td>/</td>
            <td>나눗셈</td>
            <td>정수,실수,복소수</td>
        </tr>
        <tr>
            <td>%</td>
            <td>나머지</td>
            <td>정수</td>
        </tr>
        <tr>
            <td rowspan=4>비트 연산</td>
            <td>&</td>
            <td>AND 비트 연산</td>
            <td>정수</td>
        </tr>
        <tr>
            <td>|</td>
            <td>OR 비트 연산</td>
            <td>정수</td>
        </tr>
        <tr>
            <td>^</td>
            <td>XOR 비트 연산</td>
            <td>정수</td>
        </tr>
        <tr>
            <td>&^</td>
            <td>비트 클리어</td>
            <td>정수</td>
        </tr>
         <tr>
            <td rowspan=2>시프트 연산</td>
            <td><<</td>
            <td>왼쪽 시프트</td>
            <td>정수 << 양의 정수</td>
        </tr>
        <tr>
            <td>>></td>
            <td>오른쪽 시프트</td>
            <td>정수 >> 양의 정수</td>
        </tr>
    </tbody>
</table>


**연산의 결과 타입**

Go 언어에서 모든 연산자의 각 항의 타입은 항상 같아야 합니다(시프트 연산은 예외). 예를 들어 정수 타입과 실수 타입을 서로 더하거나 뺄 수 없습니다. 그래서 타입 변환을 통해서 타입을 닽도록 맞추준 다음에 연산해야 합니다.

또 연산 결과 타입도 인수 타입과 같습니다. 즉 정수 타입과 정수 타입을 더하면 같은 정수 타입이 반환되고, 실수 타입에서 실수 티입을 나누면 실수 타입이 반환됩니다. 나머지 연산은 정수 타입만 가능합니다.



