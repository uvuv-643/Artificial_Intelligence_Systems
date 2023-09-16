import org.junit.jupiter.api.Test;
import ru.uvuv643.exceptions.ParseException;
import ru.uvuv643.io.RequestData;
import ru.uvuv643.parsers.DescriptionParser;

import static org.junit.jupiter.api.Assertions.*;

public class DescriptionParserTests {

    private final DescriptionParser descriptionParser = new DescriptionParser();

    @Test
    public void testParseMethod1() {
        RequestData data = new RequestData("Hello!;Good bye");
        assertThrows(ParseException.class, () -> descriptionParser.parse(data));
    }

    @Test
    public void testParseMethod2() {
        RequestData data = new RequestData("");
        assertThrows(ParseException.class, () -> descriptionParser.parse(data));
    }

    @Test
    public void testParseMethod3() {
        RequestData data = new RequestData("Ring     of    Basilius; Mask of Madness; Tango ");
        assertThrows(ParseException.class, () -> descriptionParser.parse(data));
    }

    @Test
    public void testParseMethod4() {
        RequestData data = new RequestData("Ring_of Basilius; Mask of Madness; Tango ");
        assertDoesNotThrow(() -> descriptionParser.parse(data));
    }

    @Test
    public void testParseMethod5() {
        RequestData data = new RequestData("tango");
        assertDoesNotThrow(() -> descriptionParser.parse(data));
    }

    @Test
    public void testParseMethod6() {
        RequestData data = new RequestData("Helm of the OVErlord Recipe");
        assertDoesNotThrow(() -> descriptionParser.parse(data));
    }

    @Test
    public void testParseMethod7() {
        RequestData data = new RequestData("Helm of the Overlord Recipe; Helm of the Overlord Recipe;");
        assertDoesNotThrow(() -> descriptionParser.parse(data));
    }

    @Test
    public void testParseMethod8() {
        RequestData data = new RequestData(";;;");
        assertDoesNotThrow(() -> descriptionParser.parse(data));
    }


}
