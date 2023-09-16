import org.junit.jupiter.api.Test;
import ru.uvuv643.helpers.OntologyHelper;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

public class OntologyHelperTests {

    @Test
    public void testGetTheMostSimilarItem1() {
        String initial = "ring_of_basilius";
        String target = "Ring_of_Basilius";
        assertEquals(OntologyHelper.getTheMostSimilarItem(initial).name, target);
    }

    @Test
    public void testGetTheMostSimilarItem2() {
        String initial = "ring_of_basiliuos";
        assertNull(OntologyHelper.getTheMostSimilarItem(initial));
    }

}
